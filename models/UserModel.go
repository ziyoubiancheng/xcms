package models

import (
	"github.com/astaxie/beego/orm"
)

type UserModel struct {
	UserId   int32  `orm:"pk;auto"`
	UserName string `orm:"size(64)"`
	AuthStr  string `orm:"size(512)"`
	Password string `orm:"size(128)"`
	IsAdmin  int8
}

func (m *UserModel) TableName() string {
	return "xcms_user"
}

func (m *UserModel) List(pageSize, page int) ([]*UserModel, int64) {
	offset := (page - 1) * pageSize

	query := orm.NewOrm().QueryTable(m.TableName())
	total, _ := query.Count()

	data := make([]*UserModel, 0)
	query.OrderBy("-user_id").Limit(pageSize, offset).All(&data)

	return data, total
}
