package models

import (
	"github.com/astaxie/beego/orm"
)

type UserModel struct {
	UserId   int    `orm:"pk;auto"`
	UserName string `orm:"size(64)"`
	AuthStr  string `orm:"size(512)"`
	Password string `orm:"size(128)"`
	IsAdmin  int8
	IsDev    int8
}

func (m *UserModel) TableName() string {
	return TbNameUser()
}

func UserList(pageSize, page int) ([]*UserModel, int64) {
	offset := (page - 1) * pageSize

	query := orm.NewOrm().QueryTable(TbNameUser())
	total, _ := query.Count()

	data := make([]*UserModel, 0)
	query.OrderBy("-user_id").Limit(pageSize, offset).All(&data)

	return data, total
}

func (m *UserModel) GetUserByName(username string) {

}
