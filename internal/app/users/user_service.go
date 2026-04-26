package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func NewLoginService() *LoginService {
	loginSvc := new(LoginService)
	return loginSvc
}

type LoginService struct {
}

func (svc *LoginService) Login(userName, password string) error {
	fmt.Println("UserName:", userName, "Password:", password)
	hashOfPassword := ""
	if !isPasswordValid(password, hashOfPassword) {
		return fmt.Errorf("password mismatch")
	}
	return nil
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func isPasswordValid(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
