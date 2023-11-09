package model

import (
	"gorm.io/gorm"
	"time"
)

type Plan struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(40);comment:plan name" json:"name"`
	Days      int            `gorm:"column:days;type:varchar(100);comment:plan days" json:"days"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
