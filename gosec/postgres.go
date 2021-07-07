package gosec

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
)

var instance *sql.DB

// GetDB returns the database single instance
func GetDB() *sql.DB {
	if instance == nil {
		log.Print("creating instance of repository")
		instance = initDb()
	}
	return instance
}

func initDb() *sql.DB {
	log.Println("connecting to db")

	connStr := "postgres://postgres:postgres@localhost:5432/db?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("could not create connection to DB, %v", err)
	}

	connPoolSize, err := strconv.Atoi(os.Getenv("CONN_POOL_SIZE"))
	if err != nil {
		panic(fmt.Errorf("could not crate the connection pool, %v", err))
	}

	db.SetMaxIdleConns(connPoolSize)
	db.SetMaxOpenConns(connPoolSize)
	db.SetConnMaxLifetime(120 * time.Minute)
	log.Println("connected to db")
	return db
}

// Close closes the database connection
func Close() {
	instance.Close()
}
