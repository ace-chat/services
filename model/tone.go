package model

import (
	"gorm.io/gorm"
	"time"
)

type Tone struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	Type      int            `gorm:"column:type;type:int(10);comment:tone type:1 content, 2 chat_bot" json:"type"`
	Name      string         `gorm:"column:name;type:varchar(20);comment:tone name" json:"name"`
	Value     string         `gorm:"column:value;type:varchar(20);comment:tone value to zubair" json:"value"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
