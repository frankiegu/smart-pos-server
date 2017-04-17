package mysql

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

/*
连接数据库
 */
func Conn(driverName, dataSourceName string) *sql.DB {

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}