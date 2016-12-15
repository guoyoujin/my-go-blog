package models

type Permissions struct {
	Id          int `orm:"pk;auto"`
	Pid         int
	Url         string
	Name        string
	Description string
	Roles       []*Roles `orm:"reverse(many)"`
}


