package user

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock/internal/utility"

	"github.com/go-playground/validator"
)

type JsonUser struct {
	Name     string `json:"name" validate:"required,min=5,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
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

var validate *validator.Validate

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

func validateUser(user *JsonUser) bool {
	validate = validator.New()
	if err := validate.Struct(user); err != nil {
		return false
	}
	return true
}

func Create(c *gin.Context) {
	var body JsonUser
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if !validateUser(&body) {
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
		Salt:      makeRandomStr(20),
		CreatedAt: time.Now().Format(utility.Layout),
		UpdatedAt: time.Now().Format(utility.Layout),
	}

	if err := utility.Db.QueryRow("SELECT TOP (1) id FROM users WHERE email = ?", user.Email).Scan(utility.TrashScanner{}); err == nil {
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
		c.Status(http.StatusCreated)
		return
	}
}

func Update(c *gin.Context) {
	var body JsonUser
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}



// func Create(name string, email string, password string) (*User, []error) {
// 	count := 50
// 	salt := makeRandomStr(20)
// 	for salt == "" && count > 0 {
// 		salt = makeRandomStr(20)
// 		count--
// 	}

// 	user := &User{
// 		Name:      name,
// 		Email:     email,
// 		Password:  password,
// 		Salt:      salt,
// 		CreatedAt: time.Now().Format(utility.Layout),
// 		UpdatedAt: time.Now().Format(utility.Layout),
// 	}

// 	validate = validator.New()

// 	if err := validate.Struct(user); err != nil {
// 		errMessages := handleValidateError(err)
// 		if errMessages != nil {
// 			return nil, errMessages
// 		}
// 	}

// 	user.Password = encrypt(user.Password + salt + os.Getenv("PEPPER"), 10000)

// 	trashScanners := make([]interface{}, 7)
// 	for i := 0; i < 7; i++ {
// 		trashScanners[i] = &utility.TrashScanner{}
// 	}
// 	if err := utility.Db.QueryRow("SELECT * FROM users WHERE email = ?", user.Email).Scan(trashScanners...); err == nil {
// 		return nil, []error{fmt.Errorf("email is already registered")}
// 	}

// 	user.Password = encrypt(user.Password, 10000)

// 	r, err := utility.Db.Exec("INSERT INTO users (name, email, password, salt, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Salt, user.CreatedAt, user.UpdatedAt)
// 	if err != nil {
// 		return nil, []error{err}
// 	}
// 	id, err := r.LastInsertId()
// 	if err != nil {
// 		return nil, []error{err}
// 	}
// 	user.Id = int(id)
// 	return user, nil
// }

// func Get(email string, password string) (*User, error) {
// 	if email == "" || password == "" {
// 		return nil, fmt.Errorf("email or password is empty")
// 	}

// 	user := &User{}
// 	if err := utility.Db.QueryRow("SELECT * FROM users WHERE email = ? LIMIT 1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt); err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, fmt.Errorf("account does not exist")
// 		}
// 		return nil, err
// 	}

// 	if user.Password != encrypt(password + user.Salt + os.Getenv("PEPPER"), 10000) {
// 		return nil, fmt.Errorf("password is not correct")
// 	}
// 	return user, nil
// }

// func Auth(email string, password string) error {
// 	if email == "" || password == "" {
// 		return fmt.Errorf("email or password is empty")
// 	}

// 	var passwordHash string
// 	var salt string
// 	if err := utility.Db.QueryRow("SELECT password, salt FROM users WHERE email = ? LIMIT 1", email).Scan(&passwordHash, &salt); err != nil {
// 		if err == sql.ErrNoRows {
// 			return fmt.Errorf("account does not exist")
// 		}
// 		return err
// 	}

// 	if passwordHash != encrypt(password + salt + os.Getenv("PEPPER"), 10000) {
// 		return fmt.Errorf("password is not correct")
// 	}
// 	return nil
// }

// func (user *User) UpdateName(name string) error {
// 	if name == "" {
// 		user.Name = ""
// 	} else {
// 		user.Name = name
// 	}
// 	user.UpdatedAt = time.Now().Format(utility.Layout)

// 	validate = validator.New()
// 	err := validate.Struct(user)
// 	if err != nil {
// 		if _, ok := err.(*validator.InvalidValidationError); ok {
// 			return err
// 		}
// 		return err
// 	}

// 	r, err := utility.Db.Exec("UPDATE users SET name = ?, updated_at = ? WHERE id = ?", user.Name, user.UpdatedAt, user.Id)
// 	if err != nil {
// 		return err
// 	}
// 	if c, err := r.RowsAffected(); err != nil || c == 0 {
// 		return fmt.Errorf("failed to update name")
// 	}
// 	return nil
// }

// func (user *User) UpdateEmail(email string) error {
// 	if email == "" {
// 		return fmt.Errorf("email is empty")
// 	}
// 	user.Email = email
// 	user.UpdatedAt = time.Now().Format(utility.Layout)

// 	err := validate.Struct(user)
// 	if err != nil {
// 		if _, ok := err.(*validator.InvalidValidationError); ok {
// 			return err
// 		}
// 		return err
// 	}

// 	r, err := utility.Db.Exec("UPDATE users SET email = ?, updated_at = ? WHERE id = ?", user.Email, user.UpdatedAt, user.Id)
// 	if err != nil {
// 		return err
// 	}
// 	if c, err := r.RowsAffected(); err != nil || c == 0 {
// 		return fmt.Errorf("failed to update email")
// 	}
// 	return nil
// }

// func (user *User) UpdatePassword(password string) error {
// 	if password == "" {
// 		return fmt.Errorf("password is empty")
// 	}
// 	salt := makeRandomStr(20)
// 	user.Password = password
// 	user.Salt = salt
// 	user.UpdatedAt = time.Now().Format(utility.Layout)

// 	err := validate.Struct(user)
// 	if err != nil {
// 		if _, ok := err.(*validator.InvalidValidationError); ok {
// 			return err
// 		}
// 		return err
// 	}

// 	user.Password = encrypt(password + user.Salt + os.Getenv("PEPPER"), 10000)

// 	r, err := utility.Db.Exec("UPDATE users SET password = ?, salt = ?, updated_at = ? WHERE id = ?", user.Password, user.Salt, user.UpdatedAt, user.Id)
// 	if err != nil {
// 		return err
// 	}
// 	if c, err := r.RowsAffected(); err != nil || c == 0 {
// 		return fmt.Errorf("failed to update password")
// 	}
// 	return nil
// }

// func (user *User) Delete() error {
// 	r, err := utility.Db.Exec("DELETE FROM users WHERE id = ?", user.Id)
// 	if err != nil {
// 		return err
// 	}
// 	if c, err := r.RowsAffected(); err != nil || c == 0 {
// 		return fmt.Errorf("failed to delete account")
// 	}

// 	return nil
// }

// func handleValidateError(err error) []error {
// 	if err == nil {
// 		return nil
// 	}
// 	if _, ok := err.(*validator.InvalidValidationError); ok {
// 		return []error{err}
// 	}
// 	var errorMessages []error
// 	for _, err := range err.(validator.ValidationErrors) {
// 		var errorMessage error
// 		fieldName := err.Field()
// 		switch fieldName {
// 			case "Name":
// 				errorTag := err.Tag()
// 				switch errorTag {
// 					case "required":
// 						errorMessage = fmt.Errorf("name is required")
// 					case "min":
// 						errorMessage = fmt.Errorf("name length must be at least 5")
// 					case "max":
// 						errorMessage = fmt.Errorf("name length must be at most 100")
// 				}
// 			case "Email":
// 				errorTag := err.Tag()
// 				switch errorTag {
// 					case "required":
// 						errorMessage = fmt.Errorf("email is required")
// 					case "email":
// 						errorMessage = fmt.Errorf("email is invalid")
// 				}
// 			case "Password":
// 				errorMessage = fmt.Errorf("password is required")
// 			case "Salt", "CreatedAt", "UpdatedAt":
// 				errorMessage = fmt.Errorf("something went wrong")
// 		}
// 		errorMessages = append(errorMessages, errorMessage)
// 	}
// 	return errorMessages
// }
