package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kairo913/tasclock/internal/domain/model"
	"github.com/kairo913/tasclock/internal/env"
	"github.com/kairo913/tasclock/internal/interfaces/database"
	"github.com/kairo913/tasclock/internal/interfaces/redis"
	"github.com/kairo913/tasclock/internal/usecase"
)

type UserController struct {
	UserInteractor    usecase.UserInteractor
	SessionInteractor usecase.SessionInteractor
}

type SessionClaims struct {
	SessionId string
	jwt.RegisteredClaims
}

var JWTSecrets [2]string

func NewUserController(c context.Context, sqlhandler database.Sqlhandler, redishandler redis.Redishandler) *UserController {
	JWTSecrets[0] = MakeRandomStr(64)
	if JWTSecrets[0] == "" {
		fmt.Println("Failed to generate JWT secret")
		<-c.Done()
	}

	ticker := time.NewTicker(time.Hour)
	go func() {
		defer ticker.Stop()
	LOOP:
		for {
			select {
			case <-c.Done():
				break LOOP
			case <-ticker.C:
				JWTSecrets[1] = JWTSecrets[0]
				JWTSecrets[0] = MakeRandomStr(64)
				if JWTSecrets[0] == "" {
					fmt.Println("Failed to generate JWT secret")
					<-c.Done()
				}
			}
		}
	}()

	return &UserController{
		UserInteractor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				Sqlhandler: sqlhandler,
			},
		},
		SessionInteractor: usecase.SessionInteractor{
			SessionRepository: &redis.SessionRepository{
				Redishandler: redishandler,
			},
		},
	}
}

func (controller *UserController) Create(c *gin.Context) {
	var body struct {
		FirstName string `json:"first_name" validate:"required,min=1,max=20"`
		LastName  string `json:"last_name" validate:"required,min=1,max=20"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	salt := MakeRandomStr(64)
	if salt == "" {
		c.Status(http.StatusInternalServerError)
		return
	}

	hash_count, err := env.GetEnvAsIntOrFallback("HASH_COUNT", 100000)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	secret_salt := env.GetEnvAsStringOrFallback("PEPPER", "")
	if secret_salt == "" {
		c.Status(http.StatusInternalServerError)
		return
	}

	user := model.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  Hash(body.Password+salt+secret_salt, hash_count),
		Salt:      salt,
	}

	id, err := controller.UserInteractor.Add(user)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user.Id = int(id)

	c.JSON(http.StatusCreated, user.LastName)
}

func (controller *UserController) Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := controller.UserInteractor.FindByEmail(body.Email)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	hash_count, err := env.GetEnvAsIntOrFallback("HASH_COUNT", 100000)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	secret_salt := env.GetEnvAsStringOrFallback("PEPPER", "")
	if secret_salt == "" {
		c.Status(http.StatusInternalServerError)
		return
	}

	if user.Password != Hash(body.Password+user.Salt+secret_salt, hash_count) {
		c.Status(http.StatusUnauthorized)
		return
	}

	expire, err := env.GetEnvAsIntOrFallback("SESSION_EXPIRE", 3600)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var sessionId string

	for i := 0; i < 10000; i++ {
		sessionId = MakeRandomStr(64)
		if sessionId == "" {
			continue
		}
		value, err := controller.SessionInteractor.SessionRepository.Get(sessionId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		if value == "" {
			break
		}
	}

	if sessionId == "" {
		c.Status(http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &SessionClaims{
		SessionId: sessionId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "tasclock",
			Audience:  []string{Hash(strconv.Itoa(user.Id), hash_count)},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString([]byte(JWTSecrets[0]))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err = controller.SessionInteractor.SessionRepository.Set(sessionId, strconv.Itoa(user.Id), time.Duration(expire)*time.Second)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, user.LastName)
}

func (controller *UserController) Logout(c *gin.Context) {

}

func (controller *UserController) Auth(t string) error {
	var token *jwt.Token
	var err error
	for _, secret := range JWTSecrets {
		token, err = jwt.ParseWithClaims(t, &SessionClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			token = nil
		} else {
			break
		}
	}
	if token == nil {
		return fmt.Errorf("invalid_token")
	}

	claims, ok := token.Claims.(*SessionClaims)
	if !ok {
		return fmt.Errorf("invalid_token")
	}
	if claims.SessionId == "" || claims.Audience[0] == "" || claims.Issuer != "tasclock" {
		return fmt.Errorf("invalid_request")
	}

	userId, err := controller.SessionInteractor.Get(claims.SessionId)
	if err != nil {
		return fmt.Errorf("internal")
	}
	if userId == "" {
		return fmt.Errorf("invalid_request")
	}

	hash_count, err := env.GetEnvAsIntOrFallback("HASH_COUNT", 100000)
	if err != nil {
		return fmt.Errorf("internal")
	}

	if Hash(userId, hash_count) != claims.Audience[0] {
		controller.SessionInteractor.FlushAll()
		JWTSecrets[0] = MakeRandomStr(64)
		if JWTSecrets[0] == "" {
			fmt.Println("Failed to generate JWT secret")
			return fmt.Errorf("secret")
		}
		return fmt.Errorf("insufficient_scope")
	}

	return nil
}
