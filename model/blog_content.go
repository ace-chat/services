package model

import (
	"gorm.io/gorm"
	"time"
)

type BlogContent struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	Type      int            `gorm:"column:type;type:int(10);comment:content type: 1 intro, 2 outline, 3 entire" json:"type"`
	AdsId     uint           `gorm:"column:ads_id;type:int(255);comment:ads id" json:"ads_id"`
	UserId    uint           `gorm:"column:user_id;type:int(255);comment:user id" json:"user_id"`
	Text      string         `gorm:"column:text;type:text;comment:content text" json:"text"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
