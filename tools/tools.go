package tools

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

type Config struct {
	Host       string
	Port       string
	LogFile    string
	dbHost     string
	dbName     string
	dbUser     string
	dbPassword string
}

type App struct {
	Config Config
	Db     *sqlx.DB
}

func GetConfig() Config {

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	return Config{
		GetEnv("HOST", ""),
		GetEnv("PORT", "8080"),
		GetEnv("LOG_FILE", "app.log"),
		GetEnv("DB_HOST", ""),
		GetEnv("DB_NAME", ""),
		GetEnv("DB_USER", ""),
		GetEnv("DB_PASSWORD", ""),
	}
}

func GetDb(config Config) *sqlx.DB {
	dsn := "host=" + config.dbHost + " user=" + config.dbUser + " dbname=" + config.dbName + " sslmode=disable password=" + config.dbPassword
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to load driver: %v", err)
	}
	return db
}

func GetApp() App {
	config := GetConfig()
	db := GetDb(config)
	return App{config, db}
}

func GetWord(cnt int, one string, two string, many string) string {
	switch cnt % 10 {
	case 1:
		if cnt%100 == 11 {
			return many
		}
		return one
	case 2:
	case 3:
	case 4:
		if cnt%100 > 11 && cnt%100 < 15 {
			return many
		}
		return two
	default:
		return many
	}

	return ""
}
