package user

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock/internal/utility"

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

func makeRandomStr(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	var result string
	for _, v := range b {
		result += string(v%byte(94) + 33)
	}
	return result
}

func encrypt(char string, count int) string {
	hash := sha256.Sum256([]byte(char))
	for i := 1; i < count; i++ {
		hash = sha256.Sum256(hash[:])
	}
	return fmt.Sprintf("%x", hash)
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

	salt := makeRandomStr(20)
	count := 10
	for salt == "" && count > 0 {
		salt = makeRandomStr(20)
		count--
	}
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

	if os.Getenv("ENV") == "dev" {
		c.IndentedJSON(http.StatusCreated, user)
		return
	}

	if os.Getenv("ENV") == "prod" {
		c.JSON(http.StatusCreated, user.Name)
		return
	}
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

	if os.Getenv("ENV") == "dev" {
		c.IndentedJSON(http.StatusOK, user)
		return
	}

	if os.Getenv("ENV") == "prod" {
		c.JSON(http.StatusOK, user.Name)
		return
	}
}