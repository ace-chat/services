package middleware

import (
	"ace/auth"
	"ace/serializer"
	"github.com/gin-gonic/gin"
	"strings"
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

		c.Set("user", claims.User)
		c.Next()
	}
}
