package token

import (
	"os"
	"time"

	"github.com/AronditFire/TODO-APP/internal/entities"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int    `json:"id"`
	Login  string `json:"login"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(user entities.User) (string, error) {
	SettedClaims := Claims{
		UserID: int(user.ID),
		Login:  user.Login,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	AccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, SettedClaims)
	AccessTokenString, err := AccessToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return AccessTokenString, err
}

func GenerateRefreshToken(user entities.User) (string, error) {
	SettedClaims := Claims{
		UserID: int(user.ID),
		Login:  user.Login,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	RefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, SettedClaims)
	RefreshTokenString, err := RefreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	return RefreshTokenString, err
}
