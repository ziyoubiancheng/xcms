package models

import (
	"github.com/astaxie/beego/orm"
)

type MenuModel struct {
	Mid    int `orm:"pk"`
	Mtype  int8
	Parent int
	Seq    int
	Name   string `orm:"size(45)"`
	Fid    int
	Role   int
}

type MenuTree struct {
	MenuModel
	Child []MenuModel
}

func (m *MenuModel) TableName() string {
	return "xcms_menu"
}

func (m *MenuModel) List() []*MenuModel {
	query := orm.NewOrm().QueryTable(m.TableName())
	data := make([]*MenuModel, 0)
	query.OrderBy("mtype", "parent", "-seq").Limit(1000).All(&data)

	return data
}
