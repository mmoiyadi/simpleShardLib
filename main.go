package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:ResAdmin14@tcp(127.0.0.1:3306)/sakila")
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) as count_actor FROM actor")

	if err == nil {
		var countActor int64
		for rows.Next() {
			rows.Scan(&countActor)
			fmt.Println(countActor)
		}
	}
}
