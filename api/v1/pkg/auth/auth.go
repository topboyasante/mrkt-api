package auth

import (
	"errors"
	"math/rand"

	"github.com/topboyasante/mrkt-api/api/v1/models"
	"github.com/topboyasante/mrkt-api/internal/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Takes the hashed password and a password, and checks if they are the same.
func VerifyPassword(hashedPW, pw string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPW), []byte(pw))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("invalid password")
		}
		return err
	}
	return nil
}

// Takes a password, hashes it, and returns it as a slice of bytes.
func HashPassword(pw string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// Generates a four digit authentication token
func GenerateAuthToken() int {
	token := rand.Intn(9000) + 1000

	return token
}

// Checks if the provided email exists in the database
func IsEmailUnique(email string) bool {
	var user models.User
	res := database.DB.Where("email = ?", email).First(&user)
	return res.Error == gorm.ErrRecordNotFound
}
