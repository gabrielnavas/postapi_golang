package database

import (
	"database/sql"
	"fmt"
	"log"
	"postapi/app/posts"

	_ "github.com/lib/pq"
)

var (
	dbUsername = "postgres"
	dbPassword = "postgres"
	dbHost     = "localhost"
	dbName     = "database"
	dbPort     = "5432"
	pgConnStr  = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbName, dbPassword)
)

type PostDB interface {
	Open() error
	Close() error
	GetDB() *sql.DB
}

type DB struct {
	db *sql.DB
}

func (d *DB) Open() error {
	pg, err := sql.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}
	log.Println("Connected to Database!")
	pg.Exec(posts.CreateSchema)
	d.db = pg
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) GetDB() *sql.DB {
	return d.db
}
