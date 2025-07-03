package model

import "time"

// SystemApi @AutoMigrate()
type SystemApi struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"size:250" json:"name"`
	SecretID  string    `gorm:"size:250" json:"secret_id"`
	SecretKey string    `gorm:"size:250" json:"secret_key"`
	Status    bool      `gorm:"default:true" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
