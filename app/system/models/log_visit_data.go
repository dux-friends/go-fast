package model

import (
	"time"
)

// LogVisitData @AutoMigrate()
type LogVisitData struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	HasType   string    `gorm:"size:250;comment:关联类型" json:"has_type"`
	HasId     uint      `gorm:"size:20;comment:关联 id" json:"has_id"`
	Driver    string    `gorm:"size:250;comment:设备" json:" driver"`
	Pv        uint      `gorm:"default:1;comment:访问量" json:"pv"`
	Uv        uint      `gorm:"default:1;comment:访客量" json:"uv"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
