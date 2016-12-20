package models

import "github.com/jinzhu/gorm"

type RolePermissions struct {
    gorm.Model
    Id            uint `gorm:"primary_key"`
    RoleId       int
    PermissionId int

}
