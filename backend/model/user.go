package model

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kairo913/tasclock/utility"

	"github.com/go-playground/validator"
)

type JsonUserSignUp struct {
	Name     string `json:"name" validate:"required,min=5,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type JsonUserSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	Salt      string
	CreatedAt string
	UpdatedAt string
}

type SessionClaims struct {
	SessionId string
	jwt.RegisteredClaims
}

func encrypt(char string, count int) string {
	hash := sha256.Sum256([]byte(char))
	for i := 1; i < count; i++ {
		hash = sha256.Sum256(hash[:])
	}
	return fmt.Sprintf("%x", hash)
}

func makeSessionId(c *gin.Context) string {
	sessionId := utility.MakeRandomStr(64)
	if sessionId == "" {
		return ""
	}

	userId, err := utility.GetSession(c, sessionId)
	if err != nil {
		return ""
	}
	if userId != "" {
		return makeSessionId(c)
	}
	return sessionId
}

func SignUp(c *gin.Context) {
	var body JsonUserSignUp
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	salt := utility.MakeRandomStr(20)
	if salt == "" {
		c.Status(http.StatusInternalServerError)
		return
	}

	user := &User{
		Name:      body.Name,
		Email:     body.Email,
		Password:  encrypt(body.Password+salt+os.Getenv("PEPPER"), 100000),
		Salt:      salt,
		CreatedAt: time.Now().Format(utility.Layout),
		UpdatedAt: time.Now().Format(utility.Layout),
	}

	if err := utility.Db.QueryRow("SELECT id FROM users WHERE email = ? LIMIT 1", user.Email).Scan(utility.TrashScanner{}); err == nil {
		c.Status(http.StatusConflict)
		return
	}

	r, err := utility.Db.Exec("INSERT INTO users (name, email, password, salt, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Salt, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user.Id = int(id)

	c.JSON(http.StatusCreated, user.Name)
}

func SignIn(c *gin.Context) {
	var body JsonUserSignIn
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := &User{}

	if err := utility.Db.QueryRow("SELECT * FROM users WHERE email = ? LIMIT 1", body.Email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	if user.Password != encrypt(body.Password+user.Salt+os.Getenv("PEPPER"), 100000) {
		c.Status(http.StatusUnauthorized)
		return
	}

	sessionId := makeSessionId(c)
	if sessionId == "" {
		c.Status(http.StatusInternalServerError)
		fmt.Println("failed to make session id")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &SessionClaims{
		SessionId: sessionId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "tasclock",
			Audience: []string{fmt.Sprint(user.Id)},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	tokenString, err := token.SignedString([]byte(utility.JWTSecrets[0]))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := utility.NewSession(c, sessionId, fmt.Sprint(user.Id)); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Header("Authorization", "Bearer " + tokenString)
	c.JSON(http.StatusOK, user.Name)
}

func UserAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, "admin")
}