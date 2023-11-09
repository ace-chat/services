package model

import (
	"gorm.io/gorm"
	"time"
)

type UserPlan struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	UserId    uint           `gorm:"column:user_id;type:int(255);comment:user id" json:"user_id"`
	PlanId    uint           `gorm:"column:plan_id;type:int(255);comment:plan id" json:"plan_id"`
	ExpiredAt time.Time      `gorm:"column:expired_at" json:"expired_at"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
