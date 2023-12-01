package model

import (
	"gorm.io/gorm"
	"time"
)

type BlogAds struct {
	Id           uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	UserId       uint           `gorm:"column:user_id;type:int(255);comment:user id" json:"user_id"`
	Type         int            `gorm:"column:type;type:int(10);comment:blog type: 1 intro, 2 outline, 3 entire" json:"type"`
	Topic        string         `gorm:"column:topic;type:varchar(200);comment:topic" json:"topic"`
	ToneId       uint           `gorm:"column:tone_id;type:int(255);comment:tone id" json:"tone_id"`
	BlogType     uint           `gorm:"column:blog_type;type:int(255);comment:blog type" json:"blog_type"`
	VoiceId      uint           `gorm:"column:voice_id;type:int(255);comment:voice id" json:"voice_id"`
	Keyword      string         `gorm:"column:keyword;type:varchar(200);comment:keyword" json:"keyword"`
	MinAge       int            `gorm:"column:min_age;type:int(20);comment:min age" json:"min_age"`
	MaxAge       int            `gorm:"column:max_age;type:int(40);comment:max age" json:"max_age"`
	WordCount    int            `gorm:"column:word_count;type:int(40);comment:word count" json:"word_count"`
	OtherDetails string         `gorm:"column:other_details;type:varchar(200);comment:other details" json:"other_details"`
	LanguageId   uint           `gorm:"column:language_id;type:int(255);comment:language id" json:"language_id"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
