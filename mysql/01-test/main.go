package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name    string
	Email   string
	Address string
	Age1    string
}

func main() {
	db, _ := sql.Open("mysql", "lcs_test:lcs_test@tcp(127.0.0.1:3306)/test1")
	fmt.Println("db:", db)

	// user1 := &User{
	// 	Name:    "LiuCS",
	// 	Email:   "123@fox.com",
	// 	Address: "湖北大学",
	// 	Age1:    "23",
	// }


	_, err := db.Exec("insert into emp (ename, sal) values (?, ?)", "zhangsan", "3000")
	if err != nil {
		log.Fatal(err)
	}

}
