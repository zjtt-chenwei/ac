package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Pet struct {
	Id      int
	Name    string
	Speci   string
	variety string
	Sex     int
	Age     int
	Intro   string
	Partner int
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Changed time.Time `orm:"auto_now_add;type(datetime)"`
	User    *User     `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Pet))
}

func (p *Pet) TableName() string {
	return "pet"
}

// 函数作用:对页面的选择结果进行筛选
func ListPet(condMap map[string]string, page int, offset int) (num int64, err error, petarr []Pet) {
	o := orm.NewOrm()
	qs := o.QueryTable("pet")
	cond := orm.NewCondition()

	if condMap["speci"] != "" {
		cond = cond.And("speci", condMap["speci"])
	}
	if condMap["variety"] != "" {
		cond = cond.And("variety", condMap["variety"])
	}
	if condMap["keywords"] != "" {
		cond = cond.Or("keywords__icontains", condMap["keywords"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset = 9
	}

	start := (page - 1) * offset
	var pets []Pet
	num, err1 := qs.Limit(offset, start).All(&pets)
	return num, err1, pets
}

// 函数作用:对页筛选结果进行计数
func CountPet(condMap map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("pet")
	cond := orm.NewCondition()

	// if condMap["now"] == "petmatch"{
	// 	if condMap["Partner"] == "0" {
	// 		cond = cond.And("")
	// 	}
	// }

	if condMap["speci"] != "" {
		cond = cond.And("speci", condMap["speci"])
	}
	if condMap["variety"] != "" {
		cond = cond.And("variety", condMap["variety"])
	}
	if condMap["keywords"] != "" {
		cond = cond.Or("keywords__icontains", condMap["keywords"])
	}

	num, _ := qs.SetCond(cond).Count()
	return num
}
