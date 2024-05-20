package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock/internal/user"
	"github.com/kairo913/tasclock/internal/utility"
)

func main() {
	utility.Init()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	r.POST("/api/user/get", func(ctx *gin.Context) {
		email := ctx.Query("email")
		password := ctx.Query("pass")
		u, err := user.Get(email, password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"result": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"user": u,
		})
	})
	r.POST("/api/user/auth", func(ctx *gin.Context) {
		email := ctx.Query("email")
		password := ctx.Query("pass")
		err := user.Auth(email, password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"result": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"result": "success",
		})
	})
	r.POST("/api/user/create", func(ctx *gin.Context) {
		name := ctx.Query("name")
		email := ctx.Query("email")
		password := ctx.Query("pass")
		u, err := user.Create(name, email, password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"result": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"user": u,
		})
	})
	r.POST("/api/user/delete", func(ctx *gin.Context) {
		email := ctx.Query("email")
		password := ctx.Query("pass")
		u, err := user.Get(email, password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"result": err.Error(),
			})
			return
		}
		err = u.Delete()
		if err != nil {
			ctx.JSON(400, gin.H{
				"result": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"result": "success",
		})
	})

	srv := &http.Server{
		Addr:    ":5000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown forced: %s", err)
	}

	log.Println("server exiting")
}
