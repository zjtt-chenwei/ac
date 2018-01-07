package models

import (
	"github.com/astaxie/beego/orm"
)

type Pet struct {
	Id			int
	Speci		string
	Type		string
	Sex			int
	Age			int
	Intro		string
	UserProfile *UserProfile `orm:"reverse(many)"` 
}

func init() {
	orm.RegisterModel(new(Pet))
}

func (p *Pet) TableName() string {
    return "pet"
}

func 
