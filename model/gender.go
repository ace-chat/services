package model

import (
	"gorm.io/gorm"
	"time"
)

type Gender struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(10);comment:plan name" json:"name"`
	Value     string         `gorm:"column:value;type:varchar(10);comment:value:1male,2female,3other" json:"value"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
