package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock/controller"
	"github.com/kairo913/tasclock/model"
	"github.com/kairo913/tasclock/utility"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	utility.Init(ctx)

	r := gin.Default()

	r.POST("/api/user/signup", model.SignUp)
	r.POST("/api/user/signin", model.SignIn)

	r.GET("/api/user/admin", controller.AuthMiddleware, model.UserAdmin)

	srv := &http.Server{
		Addr:    ":8000",
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
