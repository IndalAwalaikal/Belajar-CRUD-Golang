package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/Keuangan_db")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database connection established successfully!")

	return db
}