package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Connection struct {
	db *sql.DB
}

func NewDatabaseConnection() *Connection {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	return &Connection{
		db,
	}
}

func (c *Connection) CheckStatus() {
	if err := c.db.Ping(); err != nil {
		log.Fatalf("Database is busy: %v", err)
	}
	log.Println("Database working properly")
}

func (c *Connection) Close() {
	if err := c.db.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
	log.Println("Database connection closed")
}
