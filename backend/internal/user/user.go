package user

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/kairo913/tasclock/internal/utility"

	"github.com/go-playground/validator"
)

type User struct {
	Id        int
	Name      string `validate:"required,min=1,max=100"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=8,max=40"`
	Salt 	  string `validate:"required"`
	CreatedAt string `validate:"required"`
	UpdatedAt string `validate:"required"`
}

var validate *validator.Validate

func makeRandomStr(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("failed to generate random string")
	}

	var result string
	for _, v := range b {
		result += string(v%byte(94) + 33)
	}
	return result, nil
}

func encrypt(char string, count int) string {
	hash := sha256.Sum256([]byte(char))
	for i := 1; i < count; i++ {
		hash = sha256.Sum256(hash[:])
	}
	return fmt.Sprintf("%x", hash)
}

func Create(name string, email string, password string) (*User, error) {
	if name == "" || email == "" || password == "" {
		return nil, fmt.Errorf("name, email or password is empty")
	}

	salt, err := makeRandomStr(20)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:      name,
		Email:     email,
		Password:  password + salt + os.Getenv("PEPPER"),
		Salt:      salt,
		CreatedAt: time.Now().Format(utility.Layout),
		UpdatedAt: time.Now().Format(utility.Layout),
	}

	validate = validator.New()

	err = validate.Struct(user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}
		return nil, err
	}

	trashScanners := make([]interface{}, 7)
	for i := 0; i < 7; i++ {
		trashScanners[i] = &utility.TrashScanner{}
	}
	if err = utility.Db.QueryRow("SELECT * FROM users WHERE email = ?", user.Email).Scan(trashScanners...); err == nil {
		return nil, fmt.Errorf("email is already registered")
	}

	user.Password = encrypt(user.Password, 10000)

	r, err := utility.Db.Exec("INSERT INTO users (name, email, password, salt, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Salt, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.Id = int(id)
	return user, nil
}

func Get(email string, password string) (*User, error) {
	if email == "" || password == "" {
		return nil, fmt.Errorf("email or password is empty")
	}

	user := &User{}
	if err := utility.Db.QueryRow("SELECT * FROM users WHERE email = ? LIMIT 1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("account does not exist")
		}
		return nil, err
	}

	if user.Password != encrypt(password + user.Salt + os.Getenv("PEPPER"), 10000) {
		return nil, fmt.Errorf("password is not correct")
	}
	return user, nil
}

func Auth(email string, password string) error {
	if email == "" || password == "" {
		return fmt.Errorf("email or password is empty")
	}

	var passwordHash string
	var salt string
	if err := utility.Db.QueryRow("SELECT password, salt FROM users WHERE email = ? LIMIT 1", email).Scan(&passwordHash, &salt); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("account does not exist")
		}
		return err
	}

	if passwordHash != encrypt(password + salt + os.Getenv("PEPPER"), 10000) {
		return fmt.Errorf("password is not correct")
	}
	return nil
}

func (user *User) UpdatePassword(password string) error {
	salt, err := makeRandomStr(20)
	if err != nil {
		return err
	}
	user.Salt = salt
	user.Password = encrypt(password + user.Salt + os.Getenv("PEPPER"), 10000)
	user.UpdatedAt = time.Now().Format(utility.Layout)
	r, err := utility.Db.Exec("UPDATE users SET password = ?, salt = ?, updated_at = ? WHERE id = ?", user.Password, user.Salt, user.UpdatedAt, user.Id)
	if err != nil {
		return err
	}
	if c, err := r.RowsAffected(); err != nil || c == 0 {
		return fmt.Errorf("failed to update password")
	}
	return nil
}

func (user *User) Delete() error {
	r, err := utility.Db.Exec("DELETE FROM users WHERE id = ?", user.Id)
	if err != nil {
		return err
	}
	if c, err := r.RowsAffected(); err != nil || c == 0 {
		return fmt.Errorf("failed to delete account")
	}

	return nil
}
