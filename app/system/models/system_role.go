package model

import "gorm.io/datatypes"

// SystemRole @AutoMigrate()
type SystemRole struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	Name       string         `gorm:"size:250" json:"username"`
	Permission datatypes.JSON `json:"permission"`
	Users      []SystemUser   `gorm:"many2many:system_user_role"`
}
