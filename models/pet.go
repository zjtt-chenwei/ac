package models


import(
	"time"
	
	"github.com/astaxie/beego/orm"
	
)

type Pet struct {
	Id			int
	Name		string
	Speci		string
	Type		string
	Sex			int
	Age			int
	Intro		string
	Created		time.Time	`orm:"auto_now_add;type(datetime)"`
	Changed		time.Time	`orm:"auto_now_add;type(datetime)"`
	User  		*User 		`orm:"rel(fk)"` 
}

func init() {
	orm.RegisterModel(new(Pet))
}

func (p *Pet) TableName() string {
	return "pet"
}

func ListPet()