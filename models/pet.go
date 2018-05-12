package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Pet struct {
	Id          int
	Name        string
	Speci       string
	variety     string
	Sex         byte
	Age         int
	Intro       string
	Partner     byte
	Created     time.Time    `orm:"auto_now_add;type(datetime)"`
	Changed     time.Time    `orm:"auto_now_add;type(datetime)"`
	UserProfile *UserProfile `orm:"rel(fk)"`
	PetImg      *PetImg      `orm:"null;rel(one);on_delete(set_null)"`
}

type PetImg struct {
	Id     int
	cover  byte
	imgURL string
	Pet    *Pet `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(Pet), new(PetImg))
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
		offset = 40
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
