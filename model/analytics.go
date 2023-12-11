package model

import (
	"gorm.io/gorm"
	"time"
)

type Analytics struct {
	Id           uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	UserId       uint           `gorm:"column:user_id;type:int(255);comment:user id" json:"user_id"`
	Type         int            `gorm:"column:type;type:int(10);comment:analytics type(1simple,2deep)" json:"type"`
	BusinessDesc string         `gorm:"column:business_desc;comment:business description" json:"business_desc"`
	ProductDesc  string         `gorm:"column:product_desc;comment:product description" json:"product_desc"`
	DataDesc     string         `gorm:"column:data_desc;comment:data description" json:"data_desc"`
	ServiceId    uint           `gorm:"column:service_id;type:int(255);comment:service id" json:"service_id"`
	Content      string         `gorm:"column:content;comment:content" json:"content"`
	CreatedAt    time.Time      `gorm:"column:created_at;comment:created at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;comment:updated at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;comment:deleted at" json:"deleted_at"`
}
