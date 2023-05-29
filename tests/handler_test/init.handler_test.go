package handler_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"restaurant-management/cmd/middleware"
	routes "restaurant-management/cmd/routes"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	// Database
	db       *sql.DB
	userRepo repo.UserRepo
	authRepo repo.AuthRepo

	// Cashbin
	cashbin *casbin.Enforcer

	// Service
	homeSrv  service.HomeService
	emailSrv service.EmailService
	authSrv  service.AuthService
	userSrv  service.UserService

	// JWT
	jwt middleware.JWT
)

func init() {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "postgres:latest",
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "restaurant_management",
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor: wait.ForExec([]string{"pg_isready"}).WithPollInterval(2 * time.Second).WithExitCodeMatcher(func(exitCode int) bool {
			return exitCode == 0
		}),
	}

	sqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		panic(err)
	}

	host, err := sqlC.Host(ctx)
	if err != nil {
		panic(err)
	}

	sqlPort, err := sqlC.Ports(ctx)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, req.Env["POSTGRES_USER"], req.Env["POSTGRES_PASSWORD"], sqlPort["5432/tcp"][0].HostPort, req.Env["POSTGRES_DB"])

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = initDBSchema(db)
	if err != nil {
		panic(err)
	}

	cashbin, err = casbin.NewEnforcer("../../model_test.conf", "../../policy_test.csv")
	if err != nil {
		log.Println("Cashbin: ", err)
	}

	// Initialize Repository
	userRepo = repo.NewUserRepo(db)
	authRepo = repo.NewAuthRepo(db)

	// Initialize Services
	valiSrv := service.NewValidationService()
	crySrv := service.NewCryptoService()
	emailSrv = service.NewEmailService("", "", "", "")
	authSrv = service.NewAuthService(authRepo, "")

	// Initialize Services
	homeSrv = service.NewHomeService()
	userSrv = service.NewUserService(userRepo, authRepo, valiSrv, crySrv, authSrv, emailSrv, cashbin)

	// JWT
	jwt = middleware.Authentication(authSrv, cashbin)
}

func initDBSchema(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{
		MultiStatementEnabled: false,
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../db/migration", "postgres", driver)
	if err != nil {
		return err
	}

	return m.Up()
}

func setupApp() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	v1 := router.Group("/api/v1")

	routes.UserRoute(v1, userSrv, jwt)
	routes.HomeRoute(v1, homeSrv)

	return router
}
