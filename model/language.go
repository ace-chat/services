package model

import (
	"gorm.io/gorm"
	"time"
)

type Language struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(20);comment:language name" json:"name"`
	Iso       string         `gorm:"column:iso;type:varchar(10);comment:language code" json:"iso"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
