package models


import(
	"time"
	
	"github.com/astaxie/beego/orm"

)

type Pet struct {
	Id			int
	Speci		string
	Type		string
	Sex			int
	Age			int
	Intro		string 
	Created		time.Time
	Changed		time.Time
}

func init() {
	orm.RegisterModel(new(Pet))
}

func (p *Pet) TableName() string {
	return "pet"
}

