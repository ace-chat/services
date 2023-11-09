package model

import (
	"gorm.io/gorm"
	"time"
)

type MediaAds struct {
	Id          uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	UserId      uint           `gorm:"column:user_id;type:int(255);comment:user id" json:"user_id"`
	PlatformId  uint           `gorm:"column:platform_id;type:int(255);comment:platform id" json:"platform_id"`
	BrandName   string         `gorm:"column:brand_name;type:varchar(40);comment:brand name" json:"brand_name"`
	ServiceName string         `gorm:"column:service_name;type:varchar(40);comment:service name" json:"service_name"`
	Description string         `gorm:"column:description;type:varchar(200);comment:description" json:"description"`
	ToneId      uint           `gorm:"column:tone_id;type:int(255);comment:tone id" json:"tone_id"`
	VoiceId     uint           `gorm:"column:voice_id;type:int(255);comment:voice id" json:"voice_id"`
	Region      uint           `gorm:"column:region;type:varchar(40);comment:region" json:"region"`
	Gender      uint           `gorm:"column:gender;type:int(10);comment:gender" json:"gender"`
	MinAge      int            `gorm:"column:min_age;type:int(20);comment:min age" json:"min_age"`
	MaxAge      int            `gorm:"column:max_age;type:int(40);comment:max age" json:"max_age"`
	LanguageId  uint           `gorm:"column:language_id;type:int(255);comment:language id" json:"language_id"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
