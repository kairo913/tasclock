package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kairo913/tasclock/model"
	"github.com/kairo913/tasclock/utility"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.Header("WWW-Authenticate", "Bearer realm=\"token_required\"")
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

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
		c.Header("WWW-Authenticate", "Bearer error=\"invalid_token\"")
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	claims, ok := token.Claims.(*model.SessionClaims)
	if !ok {
		c.Header("WWW-Authenticate", "Bearer error=\"invalid_token\"")
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	if claims.SessionId == "" || claims.Audience[0] == "" || claims.Issuer != "tasclock" {
		c.Header("WWW-Authenticate", "Bearer error=\"invalid_request\"")
		c.Status(http.StatusBadRequest)
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
		c.Header("WWW-Authenticate", "Bearer error=\"invalid_request\"")
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	if userId != claims.Audience[0] {
		c.Header("WWW-Authenticate", "Bearer error=\"insufficient_scope\"")
		c.Status(http.StatusForbidden)
		c.Abort()
		utility.ResetAllSession(c)
		return
	}

	c.Next()
}
