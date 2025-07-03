package model

import (
	"time"
)

// SystemUser @AutoMigrate()
type SystemUser struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	Username  string       `gorm:"uniqueIndex;size:250" json:"username"`
	Nickname  string       `gorm:"size:250" json:"nickname"`
	Avatar    string       `gorm:"size:250" json:"avatar"`
	Password  string       `gorm:"size:250" json:"-"`
	Status    bool         `gorm:"default:true" json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Roles     []SystemRole `gorm:"many2many:system_user_role"`
}
