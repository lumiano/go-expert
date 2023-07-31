package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Database struct {
	*sql.DB
}

type ServiceInterface interface {
	connect() (*Database, error)
}

type Service struct {
	databaseService *Database
}

func NewDatabaseService() *Database {

	db, err := (&Service{}).connect()

	if err != nil {
		log.Fatal("Error on connect database:\n", err)
	}

	return db
}

func (s Service) connect() (*Database, error) {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		panic("Error getting database connection string")
	}

	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Error opening database connection:\n", err)
	}

	err = conn.Ping()

	if err != nil {
		log.Fatal("Error on ping database:", err)
	}

	log.Println("Database connected successfully!")

	return &Database{
		conn,
	}, nil

}
