package main

import (
	"database/sql"
	"fmt"
	// _ "github.com/uptrace/bun/driver/pgdriver"
	_ "github.com/go-sql-driver/mysql"
)

func Query(db *sql.DB) {
	db.Query("SELECT 1+1")
}

func main() {
	// dsn := "postgres://user:password@127.0.0.1:5432/database"
	dsn := "user:password@tcp(127.0.0.1:3306)/database"
	// db, err := sql.Open("pg", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//Test if postgres supports auto-connection-pooling:
	fmt.Println("starting 1000 queries...")
	for i := 0; i < 1000; i++ {
		go Query(db)
	}
	stats := db.Stats()
	fmt.Println("MaxOpenConnections", stats.MaxOpenConnections)
	fmt.Println("OpenConnections", stats.OpenConnections)
	fmt.Println("InUse          ", stats.InUse)
	fmt.Println("Idle           ", stats.Idle)
}
