package tools

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
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
	MediaUrl   string
	isDev      string
	logFile    string
	dbHost     string
	dbName     string
	dbUser     string
	dbPassword string
}

type App struct {
	Config Config
	Db     *sqlx.DB
	Log    *zap.Logger
}

func GetConfig() Config {

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	return Config{
		GetEnv("HOST", ""),
		GetEnv("PORT", "8080"),
		GetEnv("MEDIA_URL", ""),
		GetEnv("IS_DEV", "1"),
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

func GetLogger(config Config) *zap.Logger {
	var cfg zap.Config
	if config.isDev == "0" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	cfg.OutputPaths = []string{
		config.logFile,
	}
	logger, err := cfg.Build()
	if err != nil {
		return logger
	}
	defer logger.Sync()
	return logger
}

func GetApp() App {
	config := GetConfig()
	db := GetDb(config)
	logger := GetLogger(config)
	return App{Config: config, Db: db, Log: logger}
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

func SetCocktailOfDay(app App) {
	dayCocktailId := ""
	err := app.Db.Get(&dayCocktailId, "SELECT CAST(id AS VARCHAR) FROM tbl_cocktails ORDER BY floor(random() * 5000 + 1)::int DESC LIMIT 1")
	if err != nil {
		app.Log.Error(err.Error())
		return
	}

	_, err = app.Db.Exec("UPDATE tbl_settings SET value = $1 WHERE alias = 'day_cocktail'", dayCocktailId)
	if err != nil {
		app.Log.Error(err.Error())
		return
	}
	app.Log.Info("SET DAY COCKTAIL " + dayCocktailId)
}
