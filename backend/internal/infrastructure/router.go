package infrastructure

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock/internal/interfaces/controllers"
)

func SetUpRouter(c context.Context) *gin.Engine {
	router := gin.Default()

	sqlhandler := NewSqlhandler()
	redishandler := NewRedishandler()
	err := CreateTable(sqlhandler)
	if err != nil {
		log.Println(err.Error())
	}
	userController := controllers.NewUserController(c, sqlhandler, redishandler)
	taskController := controllers.NewTaskController(sqlhandler)

	router.POST("/user/signup", userController.Create)
	router.POST("/user/login", userController.Login)

	taskGrop := router.Group("/task", AuthMiddleware(c, userController))
	{
		taskGrop.POST("/create", taskController.Create)
	}

	return router
}

func AuthMiddleware(ctx context.Context, uc *controllers.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Header("WWW-Authenticate", "Bearer realm=\"token_required\"")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		token = strings.TrimPrefix(token, "Bearer ")

		err := uc.Auth(token)
		if err != nil {
			switch err.Error() {
			case "internal":
				c.Status(http.StatusInternalServerError)
				c.Abort()
			case "invalid_request":
				c.Header("WWW-Authenticate", "Bearer error=\"invalid_request\"")
				c.Status(http.StatusBadRequest)
				c.Abort()
			case "invalid_token":
				c.Header("WWW-Authenticate", "Bearer error=\"invalid_token\"")
				c.Status(http.StatusUnauthorized)
				c.Abort()
			case "insufficient_scope":
				c.Header("WWW-Authenticate", "Bearer error=\"insufficient_scope\"")
				c.Status(http.StatusForbidden)
				c.Abort()
			case "secret":
				c.Status(http.StatusInternalServerError)
				c.Abort()
				<-ctx.Done()
			}
		}

		c.Next()
	}
}
