package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	
)

func main() {
	db, err := sql.Open("mysql", "root:shunshun@tcp(127.0.0.1:3306)/client")
	if err != nil {
		fmt.Println("sql error:", err)
	}
	fmt.Println("db:", db)
}