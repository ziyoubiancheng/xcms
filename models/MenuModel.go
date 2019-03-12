package models

type MenuModel struct {
	BaseModel
}

type MenuItem struct {
}

func (m *MenuModel) tableName() string {
	return "xcms_menu"
}

func (m *MenuModel) List() (map[int]map[string]string, error) {
	rows, err := m.BaseModel.exec("select * from " + m.tableName() + " order by mtype,mid limit 1000")
	return rows, err
}

func (m *MenuModel) Add() {

}
