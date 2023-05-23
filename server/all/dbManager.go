package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func setupDB() {
	log.Printf("Setup DB")

	// Create a database connection string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		DB_USER,
		DB_PWD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	fmt.Println("db dsn:", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("db conn:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln("db ping:", err)
	}

	fmt.Println("db conn: up!")

	/*
		log.Printf(connectionString)

		// Open a connection to the database
		var err error
		db, err = sql.Open("mysql", connectionString)
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}

		// Set database connection options
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)

		// Check if the connection is successful
		err = db.Ping()
		if err != nil {
			log.Fatal("Failed to ping the database:", err)
		}*/
}
