package models

import "time"

type Roles struct {
	Id          int `orm:"pk;auto"`
	Name        string `orm:"unique"`
	Users       []*Users `orm:"reverse(many)"`
	Permissions []*Permissions `orm:"rel(m2m)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}
