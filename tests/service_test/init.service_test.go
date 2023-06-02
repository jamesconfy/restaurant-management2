package service_test

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/service"

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
	tableRepo repo.TableRepo
	authRepo  repo.AuthRepo
	menuRepo  repo.MenuRepo
	foodRepo  repo.FoodRepo

	// Service
	emailSrv service.EmailService
	authSrv  service.AuthService
	userSrv  service.UserService
	tableSrv service.TableService
	menuSrv  service.MenuService
	foodSrv  service.FoodService
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

	userRepo = repo.NewUserRepo(db)
	authRepo = repo.NewAuthRepo(db)
	tableRepo = repo.NewTableRepo(db)
	menuRepo = repo.NewMenuRepo(db)
	foodRepo = repo.NewFoodRepo(db)

	emailSrv = service.NewEmailService("fromEmail string", "password string", "host string", "port string")
	authSrv = service.NewAuthService(authRepo, "really_random_string")

	userSrv = service.NewUserService(userRepo, authRepo, authSrv, emailSrv)
	tableSrv = service.NewTableService(tableRepo)
	menuSrv = service.NewMenuService(menuRepo)
	foodSrv = service.NewFoodService(foodRepo, menuRepo)
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
