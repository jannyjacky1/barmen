package db

import _ "github.com/jackc/pgx/stdlib"
import "github.com/jmoiron/sqlx"
import "github.com/jannyjacky1/barmen/tools"
import "log"

type dbConfig struct {
	dbHost     string
	dbName     string
	dbUser     string
	dbPassword string
}

func Database() *sqlx.DB {

	config := dbConfig{
		tools.GetEnv("DB_HOST", ""),
		tools.GetEnv("DB_NAME", ""),
		tools.GetEnv("DB_USER", ""),
		tools.GetEnv("DB_PASSWORD", ""),
	}

	dsn := "host=" + config.dbHost + " user=" + config.dbUser + " dbname=" + config.dbName + " sslmode=disable password=" + config.dbPassword
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to load driver: %v", err)
	}
	return db
}
