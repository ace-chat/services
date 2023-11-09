package model

import (
	"gorm.io/gorm"
	"time"
)

type Voice struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	UserId    uint           `gorm:"column:user_id;type:int(255);comment:user id" json:"user_id"`
	Name      string         `gorm:"column:name;type:varchar(20);comment:voice name" json:"name"`
	Text      string         `gorm:"column:text;type:text;comment:voice text" json:"text"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
