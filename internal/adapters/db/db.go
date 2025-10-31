package postgres

import (
	"database/sql"
	"fmt"
	"katify/pkg/logger"
	"time"

	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	Database string
}

func NewPostgresClient(cfg *PostgresConfig, log logger.Logger) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)

	log.Infof("[Postgres] Attempting to connect, connStr=%s", connStr)

	db, err := sql.Open(driver, connStr)
	if err != nil {
		log.Errorf("[Postgres] sql.Open error: %v", err)
		return nil, err
	}

	for i := 0; i < 5; i++ {
		if err := db.Ping(); err != nil {
			time.Sleep(2 * time.Second)
		} else {
			log.Infof("[Postgres] Connection successful")
			return db, nil
		}
	}

	return nil, fmt.Errorf("cannot connect to database")
}
