package model

import (
	"gorm.io/gorm"
	"time"
)

type Region struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	Country   string         `gorm:"column:country;type:varchar(80);comment:country name" json:"country"`
	Iso       string         `gorm:"column:iso;type:varchar(10);comment:country iso, iso 3166-1 alpha-2" json:"iso"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
