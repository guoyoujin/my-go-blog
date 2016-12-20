package models

import (
	"github.com/jinzhu/gorm"
)

type UserRoles struct {
	gorm.Model
	User   User `gorm:"ForeignKey:user_id`
	Role   Role `gorm:"ForeignKey:role_id`
	user_id int
	role_id int
}
