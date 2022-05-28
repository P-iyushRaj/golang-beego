package models

import (
	"github.com/astaxie/beego/orm"
)

type Partner struct {
	Id       int    `orm:"auto"`
	Relation string `orm:"size(100)"`
}

func Post(partner Partner) (*Partner, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&partner)
	if err == nil {
		return nil, err
	}
	return Get(int(id))
}
func Get(id int) (*Partner, error) {
	o := orm.NewOrm()
	user := Partner{Id: id}
	err := o.Read(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
