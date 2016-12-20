package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Token     string
	Avatar    string
	Email     string
	Signature string
	Username  string
	Password  string
	Roles         []Role `gorm:"many2many:user_roles;"`
}
