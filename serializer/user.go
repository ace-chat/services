package serializer

import (
	"ace/auth"
	"ace/model"
	"time"
)

type User struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `gorm:"column:username;type:varchar(80);comment:username" json:"username"`
}

type LoginUser struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

func GeneratorUser(u model.User) Response {
	token, err := auth.GenerateToken(u.Id, u)
	if err != nil {
		return TokenError(err)
	}

	user := User{
		Id:        u.Id,
		CreatedAt: u.CreatedAt,
		Username:  u.Username,
	}

	return Response{
		Code: 200,
		Data: LoginUser{
			User:  user,
			Token: token,
		},
	}
}
