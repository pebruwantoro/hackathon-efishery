package helpers

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pebruwantoro/hackathon-efishery/config"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Email    string `json:"email"`
	Exp      int64  `json:"exp"`
	Name     string `json:"name"`
	Role     uint   `json:"role"`
	Username string `json:"username"`
	ID       uint   `json:"uuid"`
}

func GenerateJWT(ctx context.Context, user entity.User) (token string, err error) {
	cfg := config.Load()
	secret := cfg.Token.Secret
	duration, _ := time.ParseDuration(cfg.Token.Expired)

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["role"] = user.RoleId
	claims["exp"] = time.Now().Add(duration).Unix()
	jwtNew := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtNew.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
