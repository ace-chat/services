package middleware

import (
	"ace/auth"
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}
		tokenString := strings.Split(authorization, " ")
		if len(tokenString) != 2 {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}
		if tokenString[0] != "Bearer" || tokenString[1] == "" {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(tokenString[1])
		if err != nil {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}

		var user model.User
		if err := cache.DB.Model(&model.User{}).Where("id = ?", claims.User.Id).First(&user).Error; err != nil {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
