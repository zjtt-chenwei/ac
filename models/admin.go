package models

import (
	"github.com/astaxie/beego/orm"
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



// 得到某字段所有可能值（不重复）
func ShowValues(tableName string, fatherTableName string, fatherTableKey string, field string) (int64, error, []string) {
	o := orm.NewOrm()
	qs := o.QueryTable(tableName)


	var list orm.ParamsList
	var valueList []string
	var num int64
	var err error
	var fk interface{}
	if fatherTableName == "" {
		num, err = qs.ValuesFlat(&list, field)
		for _, param := range list {
			valueList = append(valueList, param.(string))
		}

	} else {
		var fatherMaps []orm.Params
		num, err = orm.NewOrm().QueryTable(fatherTableName).Filter("Name", fatherTableKey).Values(&fatherMaps)
		for _, m := range fatherMaps {
			if m["Name"] == fatherTableKey {
				fk = m["Id"]
			}
		}

		num, err = qs.Filter(fatherTableName, fk).ValuesFlat(&list, field)
		for _, param := range list {
			valueList = append(valueList, param.(string))
		}
	}
	return num, err, valueList
}
