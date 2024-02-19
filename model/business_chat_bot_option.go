package model

import (
	"gorm.io/gorm"
	"time"
)

type BusinessChatBotOption struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	Option    string         `gorm:"column:option;type:varchar(100);comment:option" json:"option"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
