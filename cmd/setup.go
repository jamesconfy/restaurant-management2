package cmd

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	middleware "restaurant-management/cmd/middleware"
	routes "restaurant-management/cmd/routes"
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

	// Validation Service
	validatorSrv := service.NewValidationService()

	// Cryptography Service
	cryptoSrv := service.NewCryptoService()

	// Home Service
	homeSrv := service.NewHomeService()

	// User Service
	userSrv := service.NewUserService(userRepo, authRepo, validatorSrv, cryptoSrv, authSrv, emailSrv, cashbin)

	// Table Service
	tableSrv := service.NewTableService(tableRepo, validatorSrv, cashbin)

	// Middleware
	jwt := middleware.Authentication(authSrv, cashbin)

	// Routes
	routes.HomeRoute(v1, homeSrv)
	routes.UserRoute(v1, userSrv, jwt)
	routes.TableRoute(v1, tableSrv, jwt)
	routes.ErrorRoute(router)

	// Documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + addr)
}

func init() {
	flag.Parse()

	addr = utils.AppConfig.ADDR
	if addr == "" {
		addr = "8000"
	}

	secret = utils.AppConfig.SECRET_KEY_TOKEN
	if secret == "" {
		log.Println("Please provide a secret key token")
	}

	mode = utils.AppConfig.MODE
	if mode == "development" {
		loadDev()
	}

	if mode == "production" {
		loadProd()
	}

	if *migrate == "true" {
		if err := utils.Migration(dsn); err != nil {
			log.Println(err)
		}
	}
}

func loadDev() {
	gin.SetMode(gin.DebugMode)

	host := utils.AppConfig.POSTGRES_HOST
	username := utils.AppConfig.POSTGRES_USERNAME
	passwd := utils.AppConfig.POSTGRES_PASSWORD
	dbname := utils.AppConfig.POSTGRES_DBNAME

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, username, passwd, dbname)
	if dsn == "" {
		log.Println("DSN cannot be empty")
	}

	email_host = utils.AppConfig.HOST
	if email_host == "" {
		log.Println("Please provide an email host name")
	}

	email_port = utils.AppConfig.PORT
	if email_port == "" {
		log.Println("Please provide an email port")
	}

	email_passwd = utils.AppConfig.PASSWD
	if email_passwd == "" {
		log.Println("Please provide an email password")
	}

	email = utils.AppConfig.EMAIL
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
