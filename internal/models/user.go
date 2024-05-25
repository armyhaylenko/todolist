package models

import (
	"errors"

	"github.com/armyhaylenko/todolist/internal/logging"
	"github.com/armyhaylenko/todolist/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	// "encoding"
)

var ErrUserNotFound = errors.New("User not found")

type User struct {
	ID           string `json:"id" redis:"id"`
	Email        string `json:"email" redis:"email"`
	PasswordHash []byte `json:"password" redis:"password"`
}

func New(email string, password string) (User, error) {
	ID := uuid.NewString()
	passwordHash, err := hashWithSalt(password)
	if err != nil {
		logging.Logger.Infow("Failed to create user", "email", email, "error", err)
		return User{}, err
	}

	return User{
		ID:           ID,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil

}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
	return err == nil
}

func (u *User) FromMap(m map[string]string) error {
	if len(m) == 0 {
		return ErrUserNotFound
	}
	*u = User{
		ID:           m["id"],
		Email:        m["email"],
		PasswordHash: utils.BytesFromString(m["password"]),
	}
	return nil
}

func hashWithSalt(password string) (passwordHash []byte, err error) {
	return bcrypt.GenerateFromPassword([]byte(password), 12)
}
