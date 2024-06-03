package utility

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRDB() {
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	checkRDBConnect(ctx, 100)
	fmt.Println("rdb connected!")
}

func checkRDBConnect(c context.Context, count uint) {
	if _, err := rdb.Ping(c).Result(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		if count == 0 {
			log.Fatal("Rdb connection failed!")
		}
		fmt.Printf("Retry connect to rdb, %d times left\n", count)
		checkRDBConnect(c, count)
	}
}

func NewSession(c *gin.Context, sessionId, userId string) error {
	if err := rdb.Set(c, sessionId, userId, time.Hour).Err(); err != nil {
		return err
	}
	return nil
}

func GetSession(c *gin.Context, sessionId string) (string, error) {
	userId, err := rdb.Get(c, sessionId).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return userId, nil
}

func DelSession(c *gin.Context, sessionId string) error {
	if err := rdb.Del(c, sessionId).Err(); err != nil {
		return err
	}
	return nil
}