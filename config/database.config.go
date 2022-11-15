package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	databaseName := os.Getenv("DB_NAME")
	//formatting

	dbURI := os.Getenv("DB_SOURCE")

	//Opens database
	db, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}

	// verifies connection to the database is still alive
	err = db.Ping()
	if err != nil {
		fmt.Println("error in pinging")
		log.Fatal(err)

	}

	log.Println("\nConnected to database:", databaseName)

	return db

}
