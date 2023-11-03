package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// User Background login user
type User struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int(255);comment:id" json:"id"`
	Username  string         `gorm:"column:username;type:varchar(40);comment:username" json:"username"`
	Password  string         `gorm:"column:password;type:varchar(80);comment:password" json:"password"`
	GoogleKey string         `gorm:"column:google_key;type:varchar(20);comment:google auth key" json:"google_key"`
	Status    string         `gorm:"column:status;type:varchar(10);comment:account status" json:"status"`
	Remark    string         `gorm:"column:remark;type:varchar(40);comment:remark" json:"remark"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:create at time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:update at time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at;comment:delete at time" json:"deleted_at"`
}

const (
	// PassWordCost
	PassWordCost = 12
	// Active active user status
	Active string = "active"
	// Inactive inactive user status
	Inactive string = "inactive"
	// Suspend suspend user status
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
