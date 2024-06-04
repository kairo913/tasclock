package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kairo913/tasclock/model"
	"github.com/kairo913/tasclock/utility"
)

func AuthMiddleware(c *gin.Context) {
	cookieKey := os.Getenv("COOKIE_KEY")
	if cookieKey == "" {
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}
	tokenString, _ := c.Cookie(cookieKey)

	var token *jwt.Token
	var err error
	for _, secret := range utility.JWTSecrets {
		token, err = jwt.ParseWithClaims(tokenString, &model.SessionClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			token = nil
		} else {
			break
		}
	}
	if token == nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	claims, ok := token.Claims.(*model.SessionClaims)
	if !ok {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	if claims.SessionId == "" || claims.Audience[0] == "" || claims.Issuer != "tasclock" {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	userId, err := utility.GetSession(c, claims.SessionId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}
	if userId == "" {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	if userId != claims.Audience[0] {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		utility.ResetAllSession(c)
		return
	}

	c.Next()
}
