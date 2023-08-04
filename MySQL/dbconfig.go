package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //注册驱动器
)

var err error

// 主数据库配置(写数据库)
var wdb *sql.DB

const (
	wDriverName = "mysql" //驱动名必须是相应的数据库

	wUsername = "root"
	wPassword = "root"
	wProtocol = "tcp"
	wAddress  = "127.0.0.1:3306"
	wDbname   = "ylb"
)

// 从数据库配置(读数据库)
var rdb *sql.DB

const (
	rDriverName = "mysql" //驱动名必须是相应的数据库

	rUsername = "root"
	rPassword = "root"
	rProtocol = "tcp"
	rAddress  = "127.0.0.2:3306"
	rDbname   = "ylb"
)
