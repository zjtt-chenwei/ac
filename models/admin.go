package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type PetSpeci struct {
	Id      int
	Name    string
	PetVari []*PetVari `orm:"reverse(many)"`
}

type PetVari struct {
	Id       int
	Name     string
	PetSpeci *PetSpeci `orm:"rel(fk)"`
}

type Province struct {
	Id   int
	Name string
	City []*City `orm:"reverse(many)"`
}

type City struct {
	Id        int
	Id_in_pro int
	Name      string
	Province  *Province `orm:"rel(fk)"`
}

var regStruct map[string]interface{}

func init() {
	orm.RegisterModel(new(PetSpeci), new(PetVari), new(Province), new(City))

	regStruct = make(map[string]interface{})
	regStruct["Petspeci"] = PetSpeci{}
	regStruct["Province"] = Province{}
}

// 得到一张表的所有内容，并且以键值对的结构存储，键：字段名；值：该字段所有值的数组
func ShowValues(tableName string) (map[string]interface{}) {
	dbname := beego.AppConfig.String("mysqldb")
	o := orm.NewOrm()
	qs := o.QueryTable(tableName)

	var vmap map[string]interface{}
	vmap = make(map[string]interface{}) //初始化

	var list orm.ParamsList
	var fieldList orm.ParamsList
	_, err1 := o.Raw("select COLUMN_NAME from information_schema.columns where table_name= ? and table_schema= ? ", tableName,dbname).ValuesFlat(&fieldList)
	if err1 !=nil{
	}

	for _,v := range fieldList{
		filed := v.(string)
		_,err2:=qs.ValuesFlat(&list,filed)
		if err2 == nil{
			vmap[filed] = list
		}
	}
	return  vmap
}
