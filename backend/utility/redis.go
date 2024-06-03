package utility

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRDB(c context.Context) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	checkRDBConnect(c, 100)
	fmt.Println("rdb connected!")
}

func checkRDBConnect(c context.Context, count uint) {
	if _, err := Rdb.Ping(c).Result(); err != nil {
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
	if err := Rdb.Set(c, sessionId, userId, time.Hour).Err(); err != nil {
		return err
	}
	return nil
}

func GetSession(c *gin.Context, sessionId string) (string, error) {
	userId, err := Rdb.Get(c, sessionId).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return userId, nil
}

func DelSession(c *gin.Context, sessionId string) error {
	if err := Rdb.Del(c, sessionId).Err(); err != nil {
		return err
	}
	return nil
}
