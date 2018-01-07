package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/gogather/com"
)

type User struct {
	Id    		int
	Account		string
	Password 	string
	UserProfile *UserProfile  `orm:"null;rel(one);on_delete(set_null)"`
	Created  	time.Time `orm:"auto_now_add;type(datetime)"`
	Changed  	time.Time `orm:"auto_now_add;type(datetime)"`
}

type UserProfile struct {
	Id			int
	Realname	string
	Account	string
	Sex			int
	Phone		string
	Email		string
	Address		string
	Hobby		string
	Birth		string
	Intro		string
	User		*User `orm:"reverse(one)"`
}


func init() {
	orm.RegisterModel(new(User), new(UserProfile))
}

func (u *User) TableName() string {
    return "user"
}

func Register(Account string, Password string) error {
	o := orm.NewOrm()
	vaild := validation.Validation{}

	pwdmd5 := com.Md5(Password)

	user := &User{Account: Account, Password: pwdmd5}			

	qs := o.QueryTable("user")
	err := qs.Filter("Account", Account).One(user)
	if err == nil {
		return err
	} 
	user.Created = time.Now()
	user.Changed = time.Now()
	user.UserProfile.Id = user.Id
	if v := vaild.Email(Account, "Email"); !v.Ok {
		user.UserProfile.Phone = Account
	}else{
		user.UserProfile.Email = Account
	}
	_, err = o.Insert(user)
	if err != nil {
		return err
	}
	return nil
}


func CheckLog(Account string, Password string) error {
	o := orm.NewOrm()
	pwdmd5 := com.Md5(Password)

	user := &User{Account: Account,Password: pwdmd5}

	qs := o.QueryTable(user)
	err := qs.Filter("Account", Account).One(user)
	
	if err != nil {
		return err
	}

	if(user.Password == pwdmd5){
		return nil
	}

	return nil
}