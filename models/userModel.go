package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserModel struct {
	ID         uint32 `json:"id" orm:"pk;column(id)`
	Name       string `json:"name" orm:"column(name)"`
	Password   string `json:"password" orm:"column(password)"`
	CreateTime string `json:"create_time" orm:"column(create_time)"`
	Phone      string `json:"phone" orm:"column(phone)"`
}

func (this *UserModel) TableName() string {
	return "user"
}

func (this *UserModel) Insert(user *UserModel) error {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	beego.Info("insert ID ", id, err)

	return err
}

func (this *UserModel) GetByName(name string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserModel))

	qs = qs.Filter("name__contains", name)

	ex := qs.Exist()

	return ex
}
