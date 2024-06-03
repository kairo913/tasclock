package utility

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

var JWTSecrets [2]string

func MakeRandomStr(length int) string {
	b := make([]byte, length)
	for i := 0; i < 10; i++ {
		if _, err := rand.Read(b); err == nil {
			break
		}
		if i == 9 {
			return ""
		}
	}

	var result string
	for _, v := range b {
		result += string(v%byte(94) + 33)
	}
	return result
}

func ResetAllSession(c context.Context) {
	if err := Rdb.FlushDB(c).Err(); err != nil {
		log.Fatal(err)
	}
	lotateJWTSecret()
	lotateJWTSecret()
}

func lotateJWTSecret() {
	JWTSecrets[1] = JWTSecrets[0]
	JWTSecrets[0] = MakeRandomStr(64)
	fmt.Println("JWT secret lotated!")
	fmt.Println(JWTSecrets)
}

func Init(c context.Context) {
	InitDB(c)
	InitRDB(c)
	JWTSecrets[0] = MakeRandomStr(64)
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		defer ticker.Stop()
	LOOP:
		for {
			select {
			case <-c.Done():
				break LOOP
			case <-ticker.C:
				lotateJWTSecret()
			}
		}

		ticker.Stop()
		fmt.Println("ticker stopped!")
	}()
}