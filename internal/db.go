package internal

import (
	"database/sql"
	"log"

	"github.com/commerce-app-demo/order-service/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(dbConfig config.DBConfig) *sql.DB {
	db, err := sql.Open(dbConfig.Driver, dbConfig.GetDSN())
	if err != nil {
		log.Fatalf("failed to open db connection error: %s", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("ping returned error: %s\nPlease check whether database is not running", err)
	}
	return db
}
