package infrastructure

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock/internal/interfaces/controllers"
)

func SetUpRouter(c context.Context) *gin.Engine {
	router := gin.Default()
	router.ContextWithFallback = true

	sqlhandler := NewSqlhandler()
	redishandler := NewRedishandler()
	err := CreateTable(sqlhandler)
	if err != nil {
		log.Println(err.Error())
	}
	userController := controllers.NewUserController(c, sqlhandler, redishandler)
	taskController := controllers.NewTaskController(sqlhandler)

	router.Use(CorsMiddleware(), ContentTypeMiddleware())

	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", userController.Create)
		userGroup.POST("/login", userController.Login)
		userGroup.POST("/logout", AuthMiddleware(c, userController), userController.Logout)
	}

	taskGroup := router.Group("/task", AuthMiddleware(c, userController))
	{
		taskGroup.POST("/create", taskController.Create)
		taskGroup.POST("/update", taskController.Update)
		taskGroup.GET("", taskController.Get)
		taskGroup.POST("/delete", taskController.Delete)
	}

	return router
}

func CorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Authorization"}
	return cors.New(config)
}

func ContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.Status(http.StatusBadRequest)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func AuthMiddleware(ctx context.Context, uc *controllers.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := uc.Auth(c)
		if err != nil {
			switch err.Error() {
			case "token_required":
				c.Header("WWW-Authenticate", "Bearer error=\"token_required\"")
				c.Status(http.StatusUnauthorized)
				c.Abort()
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
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}
