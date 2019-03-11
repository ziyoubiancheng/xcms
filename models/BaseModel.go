package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type BaseModel struct {
}

func (m *BaseModel) exec(query string, args ...interface{}) (*sql.Rows, error) {
	db, err := sql.Open("mysql", "root:135246@tcp(127.0.0.1:3307)/xcms?charset=utf8") //TODO
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	return db.Query(query, args...)
}
