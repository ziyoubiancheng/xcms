package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminModel struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (m *AdminModel) TableName() string {
	return "xcms_admin"
}

func AdminModelAdd(a *AdminModel) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func (m *AdminModel) GetByName(loginName string) (*AdminModel, error) {
	a := new(AdminModel)
	err := orm.NewOrm().QueryTable(m.TableName()).Filter("login_name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (m *AdminModel) GetList(page, pageSize int, filters ...interface{}) ([]*AdminModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*AdminModel, 0)
	query := orm.NewOrm().QueryTable(m.TableName())
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func (m *AdminModel) GetById(id int) (*AdminModel, error) {
	r := new(AdminModel)
	err := orm.NewOrm().QueryTable(m.TableName()).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (m *AdminModel) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

// func RoleAuthDelete(id int) (int64, error) {
// 	query := orm.NewOrm().QueryTable(TableName("role_auth"))
// 	return query.Filter("role_id", id).Delete()
// }

// func RoleAuthMultiAdd(ras []*RoleAuth) (n int, err error) {
// 	query := orm.NewOrm().QueryTable(TableName("role_auth"))
// 	i, _ := query.PrepareInsert()
// 	for _, ra := range ras {
// 		_, err := i.Insert(ra)
// 		if err == nil {
// 			n = n + 1
// 		}
// 	}
// 	i.Close() // 别忘记关闭 statement
// 	return n, err
// }
