package model

import (
	"gorm.io/gorm"
	"time"
)

type OptimizedAds struct {
	Id         uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	UserId     uint           `gorm:"column:user_id;type:int(255);comment:user id" json:"user_id"`
	Type       int            `gorm:"column:type;type:int(10);comment:blog type: 1 tone, 2 summarize, 3 paraphrase, 4 brand_voice, 5 target_audience" json:"type"`
	Text       string         `gorm:"column:text;comment:optimized content text" json:"text"`
	ToneId     uint           `gorm:"column:tone_id;type:int(255);comment:tone id" json:"tone_id"`
	VoiceId    uint           `gorm:"column:voice_id;type:int(255);comment:voice id" json:"voice_id"`
	Region     uint           `gorm:"column:region;type:varchar(200);comment:region" json:"region"`
	Gender     uint           `gorm:"column:gender;type:int(10);comment:gender: 0men, 1women" json:"gender"`
	MinAge     int            `gorm:"column:min_age;type:int(20);comment:min age" json:"min_age"`
	MaxAge     int            `gorm:"column:max_age;type:int(40);comment:max age" json:"max_age"`
	WordCount  int            `gorm:"column:word_count;type:int(40);comment:word count" json:"word_count"`
	LanguageId uint           `gorm:"column:language_id;type:int(255);comment:language id" json:"language_id"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
