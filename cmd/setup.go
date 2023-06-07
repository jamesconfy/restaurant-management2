package cmd

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	middleware "restaurant-management/cmd/middleware"
	routes "restaurant-management/cmd/routes"
	"restaurant-management/config"
	_ "restaurant-management/docs"
	"restaurant-management/internal/database"
	"restaurant-management/internal/logger"
	repo "restaurant-management/internal/repository"
	service "restaurant-management/internal/service"
	utils "restaurant-management/utils"

	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2"
	gin "github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	addr, mode, dsn, secret                       string
	cache                                         bool
	host, username, passwd, dbname, source        string
	email, email_passwd, email_host, email_port   string
	redis_username, redis_password, redis_address string
	db                                            *database.DB
	rdb                                           *redis.Client
	casbinEnforcer                                *casbin.Enforcer
)

var migrate = flag.Bool("m", false, "for migrations")
var casbinLoader = flag.Bool("c", false, "casbin loader")

func Setup() {
	router := gin.New()
	v1 := router.Group(utils.BasePath)
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())
	router.Use(middleware.CORS())

	defer db.Close()
	conn := db.GetConn()

	// Cache Repo
	cacheRepo := repo.NewRedisCache(rdb)

	// Auth Repository
	authRepo := repo.NewAuthRepo(conn)

	// User Repository
	userRepo := repo.NewUserRepo(conn)

	// Table Repository
	tableRepo := repo.NewTableRepo(conn)

	// Menu Repository
	menuRepo := repo.NewMenuRepo(conn)

	// Food Repository
	foodRepo := repo.NewFoodRepo(conn)

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

	// Menu Service
	menuSrv := service.NewMenuService(menuRepo)

	// Food Service
	foodSrv := service.NewFoodService(foodRepo, menuRepo)

	if cache && rdb != nil {
		userSrv = service.NewCachedUserService(userSrv, cacheRepo)
		authSrv = service.NewCachedAuthService(authSrv, cacheRepo)
		menuSrv = service.NewCachedMenuService(menuSrv, cacheRepo)
		tableSrv = service.NewCachedTableService(tableSrv, cacheRepo)
		foodSrv = service.NewCachedFoodService(foodSrv, cacheRepo)
	}

	// Routes
	routes.HomeRoute(v1, homeSrv)
	routes.UserRoute(v1, userSrv, authSrv, casbinEnforcer)
	routes.TableRoute(v1, tableSrv, authSrv, casbinEnforcer)
	routes.MenuRoute(v1, menuSrv, authSrv, casbinEnforcer)
	routes.FoodRoutes(v1, foodSrv, authSrv, casbinEnforcer)
	routes.ErrorRoute(router)

	// Documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + addr)
}

// Initialize App Variables
func init() {
	var err error

	initEnv()
	flag.Parse()

	db, err = database.New(dsn)
	if err != nil {
		log.Println("Error Connecting to DB: ", err)
	}

	opts, _ := pg.ParseURL(source)

	pgdb := pg.Connect(opts)

	adapter, err := pgadapter.NewAdapterByDB(pgdb, pgadapter.WithTableName(utils.CasbinDB))
	if err != nil {
		log.Println("Adapter is empty: ", err)
	}

	casbinEnforcer, err = casbin.NewEnforcer(utils.CasbinModel, adapter)
	if err != nil {
		log.Println("Cashbin: ", err)
	}

	rdb = database.NewRedisDB(redis_username, redis_password, redis_address)
	if rdb == nil {
		log.Println("Redis is nil, which should not be so")
	}

	switch mode {
	case "production":
		loadProd()
	default:
		loadDev()
	}

	if *migrate {
		if err := utils.Migration(dsn); err != nil {
			log.Println(err)
		}
	}

	if *casbinLoader {
		if err := initCasbinPolicy(utils.CasbinPolicy, casbinEnforcer); err != nil {
			panic(err)
		}
	}
}

// Initialize Environment Variables
func initEnv() {
	addr = config.Environment.ADDR
	cache = config.Environment.CACHE
	secret = config.Environment.SECRET_KEY_TOKEN
	redis_username = config.Environment.REDIS_USERNAME
	redis_password = config.Environment.REDIS_PASSWORD
	redis_address = config.Environment.REDIS_ADDRESS
	host = config.Environment.POSTGRES_HOST
	username = config.Environment.POSTGRES_USER
	passwd = config.Environment.POSTGRES_PASSWORD
	dbname = config.Environment.POSTGRES_DB
	email_host = config.Environment.HOST
	email_port = config.Environment.PORT
	email_passwd = config.Environment.PASSWD
	email = config.Environment.EMAIL
	mode = config.Environment.MODE
	source = config.Environment.DATABASE_SOURCE_DB

	if addr == "" {
		addr = "8000"
	}

	if secret == "" {
		log.Println("Please provide a secret key token")
	}

	if redis_username == "" {
		log.Println("REDIS USERNAME cannot be empty")
	}

	if redis_password == "" {
		log.Println("REDIS PASSWORD cannot be empty")
	}

	if redis_address == "" {
		log.Println("REDIS ADDRESS cannot be empty")
	}

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, username, passwd, dbname)
	if dsn == "" {
		log.Println("DSN cannot be empty")
	}

	if source == "" {
		log.Println("DATABASE_SOURCE_DB not provided")
	}

	if email_host == "" {
		log.Println("Please provide an email host name")
	}

	if email_port == "" {
		log.Println("Please provide an email port")
	}

	if email_passwd == "" {
		log.Println("Please provide an email password")
	}

	if email == "" {
		log.Println("Please provide an email address")
	}
}

func loadDev() {
	gin.SetMode(gin.DebugMode)
}

func loadProd() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.NewLogger())
	gin.DisableConsoleColor()
}

func initCasbinPolicy(filename string, enforcer *casbin.Enforcer) error {
	p, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	policies := strings.Split(string(p), "\n")

	go addPolicy(enforcer, policies)
	go addGroupingPolicy(enforcer, policies)

	return nil
}

func addPolicy(enforcer *casbin.Enforcer, policies []string) error {
	validPolicies := [][]string{}

	for _, policy := range policies {
		if strings.HasPrefix(policy, "p") {
			policyDetails := strings.Split(policy, ", ")
			validPolicies = append(validPolicies, policyDetails[1:])
		}
	}

	if len(validPolicies) > 0 {
		_, err := enforcer.AddPolicies(validPolicies)
		if err != nil {
			return err
		}
	}

	return nil
}

func addGroupingPolicy(enforcer *casbin.Enforcer, policies []string) error {
	validGrouping := [][]string{}
	for _, policy := range policies {
		if strings.HasPrefix(policy, "g") {
			policyDetails := strings.Split(policy, ", ")
			validGrouping = append(validGrouping, policyDetails[1:])
		}
	}

	if len(validGrouping) > 0 {
		_, err := enforcer.AddGroupingPolicies(validGrouping)
		if err != nil {
			return err
		}
	}

	return nil
}
