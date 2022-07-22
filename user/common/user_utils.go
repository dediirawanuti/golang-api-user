package common

import (
	"github.com/golang-api-user/user/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

// LogError use logging error
func LogError(function string, str string, err error) {
	log.Printf("[%v] %v : %v", function, str, err)
}

// Message wrapper a message response
func Message(status bool, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message, "data": data}
}

// Response for print a response
func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

// Hash use for hashing password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword use for validasi password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Encrypt use for encypt
func Encrypt(user *models.User) (*models.User, error) {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return user, nil
}

// Validate use for validate register, login, and update
func Validate(action string, user *models.User) error {
	switch strings.ToLower(action) {
	case "register":
		if user.Name == "" {
			return errors.New("Required Name")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.NoTlpn == "" {
			return errors.New("Required Phone Number")
		}
		if user.Gender == "" {
			return errors.New("Required Gender")
		}
		if user.Address == "" {
			return errors.New("Required Address")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid email address")
		}
		return nil
	case "login":
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "update":
		if user.Name == "" {
			return errors.New("Required Name")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if user.NoTlpn == "" {
			return errors.New("Required Phone Number")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Gender == "" {
			return errors.New("Required Gender")
		}
		if user.Address == "" {
			return errors.New("Required Address")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	default:
		if user.Name == "" {
			return errors.New("Required Name")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if user.NoTlpn == "" {
			return errors.New("Required Phone Number")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Gender == "" {
			return errors.New("Required Gender")
		}
		if user.Address == "" {
			return errors.New("Required Address")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}
