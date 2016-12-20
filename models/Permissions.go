package models

import "github.com/jinzhu/gorm"

type Permissions struct {
	gorm.Model
	Id           uint `gorm:"primary_key"`
	Pid         int
	Url         string
	Name        string
	Description string
}


