package database

import (
	"database/sql"
	"log"
	"postapi/app/posts"

	_ "github.com/lib/pq"
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
	log.Println("Creating tables")
	createTables(pg)
	d.db = pg
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) GetDB() *sql.DB {
	return d.db
}

func createTables(sql *sql.DB) {
	sql.Exec(posts.CreateSchema)
}
