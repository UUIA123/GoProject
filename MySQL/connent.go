package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 创建数据库连接
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ylb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the database!")
	fmt.Println("Successfully connected to the database!111111")
	fmt.Println("Successfully connected to the database!1111112222")
}
