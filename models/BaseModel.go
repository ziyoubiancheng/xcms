package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type BaseModel struct {
}

//Base SQL opt
func (m *BaseModel) exec(query string, args ...interface{}) (map[int]map[string]string, error) {
	db, err := sql.Open("mysql", "root:135246@tcp(127.0.0.1:3306)/xcms?charset=utf8") //TODO
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	row, err := db.Query(query, args...)
	results := m.readRows(row)
	return results, err
}

func (m *BaseModel) readRows(query *sql.Rows) map[int]map[string]string {
	column, _ := query.Columns()              //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return nil
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	return results
}
