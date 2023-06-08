package handler_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	routes "restaurant-management/cmd/routes"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/service"
	"restaurant-management/utils"

	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"

	"github.com/golang-migrate/migrate/v4"
	postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	// Database
	db        *sql.DB
	userRepo  repo.UserRepo
	authRepo  repo.AuthRepo
	menuRepo  repo.MenuRepo
	foodRepo  repo.FoodRepo
	tableRepo repo.TableRepo

	//
	router *gin.Engine

	// Cashbin
	cashbinEnforcer *casbin.Enforcer

	// Service
	homeSrv  service.HomeService
	emailSrv service.EmailService
	authSrv  service.AuthService
	userSrv  service.UserService
	menuSrv  service.MenuService
	foodSrv  service.FoodService
	tableSrv service.TableService
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

	username := req.Env["POSTGRES_USER"]
	password := req.Env["POSTGRES_PASSWORD"]
	dbname := req.Env["POSTGRES_DB"]
	port := sqlPort["5432/tcp"][0].HostPort

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, username, password, port, dbname)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = initDBSchema(db)
	if err != nil {
		panic(err)
	}

	source := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)
	opts, _ := pg.ParseURL(source)

	pgdb := pg.Connect(opts)

	adapter, err := pgadapter.NewAdapterByDB(pgdb, pgadapter.WithTableName(utils.CasbinDB))
	if err != nil {
		log.Println("Adapter is empty: ", err)
	}

	cashbinEnforcer, err = casbin.NewEnforcer(utils.CasbinTestModel, adapter)
	if err != nil {
		log.Println("Cashbin: ", err)
	}

	err = initCasbinPolicy(utils.CasbinTestPolicy, cashbinEnforcer)
	if err != nil {
		panic(err)
	}

	// Initialize Repository
	userRepo = repo.NewUserRepo(db)
	authRepo = repo.NewAuthRepo(db)
	menuRepo = repo.NewMenuRepo(db)
	foodRepo = repo.NewFoodRepo(db)
	tableRepo = repo.NewTableRepo(db)

	// Initialize Services
	emailSrv = service.NewEmailService("", "", "", "")
	authSrv = service.NewAuthService(authRepo, "")

	// Initialize Services
	homeSrv = service.NewHomeService()
	userSrv = service.NewUserService(userRepo, authRepo, authSrv, emailSrv)
	menuSrv = service.NewMenuService(menuRepo)
	foodSrv = service.NewFoodService(foodRepo, menuRepo)
	tableSrv = service.NewTableService(tableRepo)

	// Router
	router = setupApp()
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

	err = m.Up()
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	return nil
}

func initCasbinPolicy(filename string, enforcer *casbin.Enforcer) error {
	p, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	validPolicies := [][]string{}
	policies := strings.Split(string(p), "\n")
	for _, policy := range policies {
		if strings.HasPrefix(policy, "p") {
			policyDetails := strings.Split(policy, ", ")
			validPolicies = append(validPolicies, policyDetails[1:])
		}
	}

	_, err = enforcer.AddPolicies(validPolicies)
	return err
}

func setupApp() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	v1 := router.Group(utils.BasePath)

	routes.UserRoute(v1, userSrv, authSrv, cashbinEnforcer)
	routes.MenuRoute(v1, menuSrv, authSrv, cashbinEnforcer)
	routes.FoodRoutes(v1, foodSrv, authSrv, cashbinEnforcer)
	routes.TableRoute(v1, tableSrv, authSrv, cashbinEnforcer)
	routes.HomeRoute(v1, homeSrv)
	routes.ErrorRoute(router)

	return router
}
