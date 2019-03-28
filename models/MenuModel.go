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

func (m *MenuModel) Tree() map[int]MenuTree {
	query := orm.NewOrm().QueryTable(m.TableName())
	data := make([]*MenuModel, 0)
	query.OrderBy("mtype", "parent", "-seq").Limit(1000).All(&data)

	var menu = make(map[int]MenuTree)
	for _, v := range data { //查询出来的数组
		//fmt.Println(v.Mid, v.Parent, v.Name)
		if 0 == v.Mtype {
			var tree = new(MenuTree)
			tree.MenuModel = *v
			menu[v.Mid] = *tree
		} else {
			if tmp, ok := menu[v.Parent]; ok {
				tmp.Child = append(tmp.Child, *v)
				menu[v.Parent] = tmp
			}
		}
	}

	return menu
}

func (m *MenuModel) List() []*MenuModel {
	query := orm.NewOrm().QueryTable(m.TableName())
	data := make([]*MenuModel, 0)
	query.OrderBy("mtype", "parent", "-seq").Limit(1000).All(&data)

	return data
}
