package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"my-go-blog/utils"
)

type Users struct {
	Id        int `orm:"pk;auto"`
	Name  string `orm:"unique"`
	Token     string `orm:"unique"`
	Avatar    string
	Email     string `orm:"null"`
	Signature string `orm:"null;size(1000)"`
	Username  string `orm:"unique"`
	Password  string
	Roles     []*Roles `orm:"rel(m2m)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime)"`

}

func FindUserById(id int) (bool, Users) {
	o := orm.NewOrm()
	var user Users
	err := o.QueryTable(user).Filter("Id", id).One(&user)
	return err != orm.ErrNoRows, user
}


func Login(name string, password string) (bool, Users) {
	o := orm.NewOrm()
	var user Users
	err := o.QueryTable(user).Filter("Name", name).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByName(name string) (bool, Users) {
	o := orm.NewOrm()
	var user Users
	err := o.QueryTable(user).Filter("Name", name).One(&user)
	return err != orm.ErrNoRows, user
}

func SaveUser(user *Users) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

func UpdateUser(user *Users) {
	o := orm.NewOrm()
	o.Update(user)
}

func PageUser(p int, size int) utils.Page {
	o := orm.NewOrm()
	var user Users
	var list []Users
	qs := o.QueryTable(user)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-CreatedAt").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

func DeleteUser(user *Users) {
	o := orm.NewOrm()
	o.Delete(user)
}
