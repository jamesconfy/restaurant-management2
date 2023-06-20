package repo_test

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	repo "restaurant-management/internal/repository"

	"github.com/golang-migrate/migrate/v4"
	postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	db            *sql.DB
	userRepo      repo.UserRepo
	authRepo      repo.AuthRepo
	tableRepo     repo.TableRepo
	foodRepo      repo.FoodRepo
	menuRepo      repo.MenuRepo
	orderRepo     repo.OrderRepo
	orderItemRepo repo.OrderItemRepo
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
	foodRepo = repo.NewFoodRepo(db)
	menuRepo = repo.NewMenuRepo(db)
	orderRepo = repo.NewOrderRepo(db)
	orderItemRepo = repo.NewOrderItemRepo(db)
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
