package models

import (
	"fmt"
)

type MenuModel struct {
	BaseModel
}

func (m *MenuModel) TableName() string {
	return "xcms_menu"
}

func (m *MenuModel) List() {
	rows, err := m.BaseModel.exec("select mid,name from xcms_menu limit 10")
	mid := 0
	name := " "
	for rows.Next() {
		rows.Scan(&mid, &name)
		fmt.Println(mid, name)
	}
	fmt.Println(err)
}
