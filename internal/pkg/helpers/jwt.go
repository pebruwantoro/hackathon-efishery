package helpers

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pebruwantoro/hackathon-efishery/config"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func GenerateJWT(ctx context.Context, user entity.User) (token string, err error) {
	cfg := config.Load()
	secret := cfg.Token.Secret
	duration, _ := time.ParseDuration(cfg.Token.Expired)

	claims := jwt.MapClaims{}
	claims["uuid"] = user.UUID
	claims["name"] = user.Name
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["role"] = user.AccessRole
	claims["exp"] = time.Now().Add(duration)
	jwtNew := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtNew.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
