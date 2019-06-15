package models

import (
	"github.com/astaxie/beego/orm"
)

type DataModel struct {
	Did        int    `orm:"pk;auto"`
	Mid        int    `orm:"default(0)"`
	Parent     int    `orm:"default(0)"`
	Name       string `orm:"size(60)"`
	Content    string `orm:"size(2048);default({})"`
	Seq        int    `orm:"index"`
	Status     int8
	UpdateTime int64
}

type DataStruct struct {
	DataModel
	ContentStruct interface{}
}

func (c *DataModel) TableName() string {
	return TbNameData()
}

func DataList(mid, pageSize, page int) ([]*DataModel, int64) {
	if mid <= 0 {
		return nil, 0
	}

	offset := (page - 1) * pageSize
	query := orm.NewOrm().QueryTable(TbNameData()).Filter("mid", mid)
	total, _ := query.Count()
	data := make([]*DataModel, 0)
	query.OrderBy("parent", "-seq").Limit(pageSize, offset).All(&data)

	return data, total
}
