package hash

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AppBcrypt interface {
	HashAndSalt(pwd []byte) (string, error)
	ComparePasswords(hashedPwd string, plainPwd []byte) bool
}

type AppBcryptImpl struct {
}

// HashAndSalt Hashes a given string
func (d AppBcryptImpl) HashAndSalt(pwd []byte) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", errors.New(fmt.Sprintf("An unexpected error ocurred while hashing the password: %v", err))
	}
	return string(hash), nil
}

func (d AppBcryptImpl) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
