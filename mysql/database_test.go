package mysql

import (
	"testing"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestConn(t *testing.T) {

	db := Conn("mysql", "test:cardinfolink@tcp(192.168.1.177:3306)/msp?charset=utf8")

	rows, err := db.Query("select * from term_info where term_id = 139213");

	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}
