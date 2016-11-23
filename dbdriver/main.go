package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@/gosample")
	errorChecker(err)
	fmt.Println(db)
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
