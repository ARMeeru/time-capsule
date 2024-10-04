package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/ARMeeru/time-capsule/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		// First, try to get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// If not in header, try to get it from the cookie
			tokenCookie, err := c.Cookie("token")
			if err == nil {
				tokenString = tokenCookie
			}
		}

		if tokenString == "" {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return []byte(config.GetEnv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", uint(claims["user_id"].(float64)))

		c.Next()
	}
}
