package cmd

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	middleware "restaurant-management/cmd/middleware"
	routes "restaurant-management/cmd/routes"
	"restaurant-management/config"
	_ "restaurant-management/docs"
	sql "restaurant-management/internal/database"
	"restaurant-management/internal/logger"
	repo "restaurant-management/internal/repository"
	service "restaurant-management/internal/service"
	utils "restaurant-management/utils"

	"github.com/casbin/casbin/v2"
	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	addr         string
	mode         string
	dsn          string
	secret       string
	email        string
	email_passwd string
	email_host   string
	email_port   string
)

var migrate = flag.String("migrate", "false", "for migrations")

func Setup() {
	router := gin.New()
	v1 := router.Group("/api/v1")
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())
	router.Use(middleware.CORS())

	db, err := sql.New(dsn)
	if err != nil {
		log.Println("Error Connecting to DB: ", err)
	}
	defer db.Close()
	conn := db.GetConn()

	cashbin, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Println("Cashbin: ", err)
	}

	// Auth Repository
	authRepo := repo.NewAuthRepo(conn)

	// User Repository
	userRepo := repo.NewUserRepo(conn)

	// Table Repository
	tableRepo := repo.NewTableRepo(conn)

	// Email Service
	emailSrv := service.NewEmailService(email, email_passwd, email_host, email_port)

	// Token Service
	authSrv := service.NewAuthService(authRepo, secret)

	// Home Service
	homeSrv := service.NewHomeService()

	// User Service
	userSrv := service.NewUserService(userRepo, authRepo, authSrv, emailSrv)

	// Table Service
	tableSrv := service.NewTableService(tableRepo)

	// Routes
	routes.HomeRoute(v1, homeSrv)
	routes.UserRoute(v1, userSrv, authSrv, cashbin)
	routes.TableRoute(v1, tableSrv, authSrv, cashbin)
	routes.ErrorRoute(router)

	// Documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + addr)
}

func init() {
	flag.Parse()

	addr = config.Environment.ADDR
	if addr == "" {
		addr = "8000"
	}

	secret = config.Environment.SECRET_KEY_TOKEN
	if secret == "" {
		log.Println("Please provide a secret key token")
	}

	mode = config.Environment.MODE
	switch mode {
	case "production":
		loadProd()
	default:
		loadDev()
	}

	if *migrate == "true" {
		if err := utils.Migration(dsn); err != nil {
			log.Println(err)
		}
	}
}

func loadDev() {
	gin.SetMode(gin.DebugMode)

	host := config.Environment.POSTGRES_HOST
	username := config.Environment.POSTGRES_USER
	passwd := config.Environment.POSTGRES_PASSWORD
	dbname := config.Environment.POSTGRES_DB

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, username, passwd, dbname)
	fmt.Println("DSN: ", dsn)
	if dsn == "" {
		log.Println("DSN cannot be empty")
	}

	email_host = config.Environment.HOST
	if email_host == "" {
		log.Println("Please provide an email host name")
	}

	email_port = config.Environment.PORT
	if email_port == "" {
		log.Println("Please provide an email port")
	}

	email_passwd = config.Environment.PASSWD
	if email_passwd == "" {
		log.Println("Please provide an email password")
	}

	email = config.Environment.EMAIL
	if email == "" {
		log.Println("Please provide an email address")
	}
}

func loadProd() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.NewLogger())
	gin.DisableConsoleColor()
}

var _ = loadProd
