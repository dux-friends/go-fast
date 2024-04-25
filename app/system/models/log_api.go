package model

import (
	"gorm.io/datatypes"
	"time"
)

// LogApi @AutoMigrate()
type LogApi struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:250;comment:路由名" json:"name"`
	Method    string         `gorm:"size:20;comment:请求方法" json:"method"`
	HasType   string         `gorm:"size:20;comment:关联类型" json:"has_type"`
	Date      datatypes.Date `gorm:"comment:日期" json:"date"`
	Pv        uint           `gorm:"default:1;comment:访问量" json:"pv"`
	Uv        uint           `gorm:"default:1;comment:访客量" json:"uv"`
	MaxTime   float64        `gorm:"precision:3;comment:访客量" json:"max_time"`
	MinTime   float64        `gorm:"precision:3;comment:访客量" json:"min_time"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
