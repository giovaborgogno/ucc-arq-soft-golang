package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"crud_app/config"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbHost, cfg.DbPort)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
