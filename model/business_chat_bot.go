package model

import (
	"gorm.io/gorm"
	"time"
)

type BusinessChatBot struct {
	Id                  uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	UserId              uint           `gorm:"column:user_id;type:int(255);comment:user id" json:"user_id"`
	CompanyName         string         `gorm:"column:company_name;type:varchar(100);comment:company name" json:"company_name"`
	CompanyIntroduction string         `gorm:"column:company_introduction" json:"company_introduction"`
	PhoneNumber         string         `gorm:"column:phone_number;type:varchar(100);comment:phone number" json:"phone_number"`
	ToneId              uint           `gorm:"column:tone_id;type:int(255);comment:tone id" json:"tone_id"`
	CreatedAt           time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
