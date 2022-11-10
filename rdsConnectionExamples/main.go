package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/sijms/go-ora/v2"
)

//POSTGRESQL
// func main() {
// 	db, err := sql.Open("pgx", os.Getenv("PG_DSN"))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer db.Close()

// 	err = db.Ping()

// 	fmt.Println(err)
// }

// MYSQL

// func main() {
// 	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")
// }

// MARIADB
// var db *sql.DB

// func main() {
// 	// Get a database handle.
// 	var err error
// 	db, err = sql.Open("mysql", os.Getenv("MARIADB_DSN"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")
// }

// SQL SERVER

// func main() {
// 	db, err := sql.Open("sqlserver", os.Getenv("MSSQL_DSN"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")
// }

//ORACLE

func main() {
	db, err := sql.Open("oracle", os.Getenv("ORACLE_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
