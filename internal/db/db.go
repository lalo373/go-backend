package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func Connect(dburl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dburl)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)

	context.WithTimeout(context.Background(), 5 * time.Second)
 	return db, nil
}
