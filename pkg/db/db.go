package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	Client *sql.DB
}

func Get(connectionString string) (*DB, error) {
	db, err := get(connectionString)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}

func get(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Check if we can ping
	if err := db.Ping(); err != nil {
		return nil, err
	} else {
		log.Println("Connected to database")
	}

	return db, nil
}
