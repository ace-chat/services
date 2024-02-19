package model

import (
	"gorm.io/gorm"
	"time"
)

type BusinessChatBotQA struct {
	Id            uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	BusinessBotId uint           `gorm:"column:business_bot_id;type:int(255);comment:business bot id" json:"business_bot_id"`
	Question      string         `gorm:"column:question;comment:question text" json:"question"`
	Answer        string         `gorm:"column:answer;comment:answer text" json:"answer"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
