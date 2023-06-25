package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alvinatthariq/farmsvc-go/controllers"
	"github.com/alvinatthariq/farmsvc-go/domain"
	"github.com/alvinatthariq/farmsvc-go/entity"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dbgorm      *gorm.DB
	router      *mux.Router
	redisClient *redis.Client
	err         error

	dom domain.DomainItf
)

func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database SQL
	ConnectSQL(AppConfig.MySQL.ConnectionString)
	MigrateSQL()

	// Initialize Redis
	ConnectRedis()

	// Initialize the router
	router = mux.NewRouter().StrictSlash(true)

	// Initialize domain
	dom = domain.Init(dbgorm, redisClient)

	// Initialize controller
	controllers.Init(dbgorm, router, dom)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func ConnectSQL(connectionString string) {
	dbgorm, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func MigrateSQL() {
	dbgorm.AutoMigrate(&entity.Farm{})
	dbgorm.AutoMigrate(&entity.Pond{})
	log.Println("Database Migration Completed...")
}

func ConnectRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Host,
		Password: AppConfig.Redis.Password,
	})
}
