package models

import "github.com/astaxie/beego/orm"

type PetSpeci struct {
	Id      int
	Speci   string
	PetVari []*PetVari `orm:"reverse(many)"`
}

type PetVari struct {
	Id       int
	Variety  string
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

func init() {
	orm.RegisterModel(new(PetSpeci), new(PetVari), new(Province), new(City))
}

// 得到某字段所有可能值（不重复）
func DistinctNum(tableName string, fatherTablekey string, s string) (int64, error, []string) {
	o := orm.NewOrm()
	qs := o.QueryTable(tableName)

	var list orm.ParamsList
	var typelist []string
	var num int64
	var err error
	if fatherTablekey != "" {
		num, err = qs.Distinct().ValuesFlat(&list, s)
		for _, param := range list {
			typelist = append(typelist, param.(string))
		}
	}

	return num, err, typelist
}
