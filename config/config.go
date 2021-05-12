package config

import (
	"crypto/tls"
	"database/sql"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	logo "log"

	"github.com/go-redis/redis/v7"
	"github.com/jasonbronson/go-gin-boilerplate/library/log"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
	_ "github.com/newrelic/go-agent/v3/integrations/nrsqlite3"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/xo/dburl"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Cfg    = &Config{}
	Driver = "nrsqlite3"
)

type Config struct {
	Port               int
	DatabaseURL        string
	DBLogMode          bool
	RedisDB            int
	RedisClient        *redis.Client
	JWTSecret          string
	JWTIssuer          string
	JWTAudience        string
	SQLDB              *sql.DB
	GormDB             *gorm.DB
	NewRelicEnabled    bool
	NewRelicLicenseKey string
	NewRelicAppName    string
	NewRelicApp        *newrelic.Application
}

func init() {
	initEnv()
	initDB()
	initRedis()
}

func initEnv() {
	godotenv.Load()
	Cfg.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	Cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	Cfg.DBLogMode, _ = strconv.ParseBool(os.Getenv("DB_LOG_MODE"))
	Cfg.RedisDB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	Cfg.JWTSecret = os.Getenv("JWT_SECRET")
	Cfg.JWTIssuer = os.Getenv("JWT_ISSUER")
	Cfg.JWTAudience = os.Getenv("JWT_AUDIENCE")
	Cfg.NewRelicEnabled, _ = strconv.ParseBool(os.Getenv("NEW_RELIC_ENABLED"))
	Cfg.NewRelicLicenseKey = os.Getenv("NEW_RELIC_LICENSE_KEY")
	Cfg.NewRelicAppName = os.Getenv("NEW_RELIC_APP_NAME")

}

func initDB() {

	var err error
	u, err := dburl.Parse(Cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	dbname := u.DSN
	if dbname == "" {
		log.Fatal("database not found or empty env var")
	}
	dbdialect := sqlite.Open(dbname)

	newLogger := logger.New(
		logo.New(logo.Writer(), "GORMDB:", logo.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  logger.Error,    // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Enable color
		},
	)

	gconfig := &gorm.Config{}
	if Cfg.GormDB, err = gorm.Open(dbdialect, gconfig); err != nil {
		log.Fatalf("could not initialize gorm: %v", err)
	}

	//Debug SQL logs?
	if Cfg.DBLogMode {
		Cfg.GormDB.Logger = newLogger
	}

	log.Println("Success connecting to database")
}

func initRedis() {
	redisPool, err := strconv.ParseInt(os.Getenv("REDIS_POOL_SIZE"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	redisURL, _ := url.Parse(os.Getenv("REDIS_URL"))
	redisPassword, _ := redisURL.User.Password()
	redisHost := redisURL.Host
	redisOptions := &redis.Options{
		Addr:       redisHost,
		Password:   redisPassword,
		DB:         Cfg.RedisDB,
		PoolSize:   int(redisPool),
		MaxRetries: 2,
	}

	if strings.Contains(os.Getenv("REDIS_URL"), "rediss") {
		redisOptions.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	Cfg.RedisClient = redis.NewClient(redisOptions)
	Cfg.RedisClient.AddHook(nrredis.NewHook(redisOptions))

	pong, err := Cfg.RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("Redis cannot connect ", err)
	}
	log.Println("Redis pong:", pong)
}
