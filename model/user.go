package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// User Background login user
type User struct {
	Id          uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	Username    string         `gorm:"column:username;type:varchar(40);comment:username" json:"username"`
	DisplayName string         `gorm:"column:display_name;type:varchar(80);comment:display name" json:"display_name"`
	Email       string         `gorm:"column:email;type:varchar(120);comment:email" json:"email"`
	Avatar      string         `gorm:"column:avatar;type:varchar(120);comment:user avatar" json:"avatar"`
	Phone       string         `gorm:"column:phone;type:varchar(100);comment:user phone" json:"phone"`
	Password    string         `gorm:"column:password;type:varchar(200);comment:password" json:"password"`
	PlanId      uint           `gorm:"column:plan_id;type:int(255);comment:user plan, according to plan table" json:"plan_id"`
	Status      string         `gorm:"column:status;type:varchar(10);comment:" json:"status"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

const (
	// PassWordCost
	PassWordCost = 12
	// Active user status
	Active string = "active"
	// Inactive user status
	Inactive string = "inactive"
	// Suspend user status
	Suspend string = "suspend"
)

// SetPassword set encode password
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	user.Status = Active
	return nil
}

// CheckPassword verify encode password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
