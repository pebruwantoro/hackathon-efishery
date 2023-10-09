package helpers

import (
	"fmt"

	"github.com/pebruwantoro/hackathon-efishery/config"
	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	cfg := config.Load()
	passwordSalt := fmt.Sprintf("%s%s", password, cfg.App.Salt)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	cfg := config.Load()
	passwordSalt := fmt.Sprintf("%s%s", password, cfg.App.Salt)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordSalt))
	if err != nil {
		return err
	}
	return nil
}
