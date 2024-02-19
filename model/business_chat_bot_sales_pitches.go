package model

import (
	"gorm.io/gorm"
	"time"
)

type BusinessChatBotSalesPitch struct {
	Id            uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	BusinessBotId uint           `gorm:"column:business_bot_id;type:int(255);comment:business bot id" json:"business_bot_id"`
	Topic         string         `gorm:"column:topic;comment:topic text" json:"topic"`
	Input         string         `gorm:"column:input;comment:input text" json:"input"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
