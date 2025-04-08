package repositories

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/isjhar/iet/internal/config"
	"github.com/isjhar/iet/internal/domain/entities"
)

const dataKey string = "data"

type JwtRepository struct {
}

func (r JwtRepository) GenerateToken(data interface{}) (string, error) {
	plainToken := jwt.New(jwt.SigningMethodHS512)
	claims := plainToken.Claims.(jwt.MapClaims)
	claims[dataKey] = data
	claims["exp"] = time.Now().Add(time.Second * time.Duration(config.Jwt.AccessTokenExpiresIn.Int64)).Unix()

	securedToken, err := plainToken.SignedString([]byte(r.GetJwtSecret()))
	if err != nil {
		return "", entities.InternalServerError
	}

	return securedToken, nil
}

func (r JwtRepository) GetData(token string) (interface{}, error) {
	claims, err := r.getClaims(token)
	if err != nil {
		return nil, err
	}
	data, ok := claims[dataKey]
	// If the key exists
	if !ok {
		return nil, entities.EntityNotFound
	}
	return data, nil
}

func (r JwtRepository) getClaims(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.GetJwtSecret()), nil
	})
	if err != nil {
		return nil, entities.ErrorUnauthorized
	}
	return claims, nil
}

func (r JwtRepository) GetJwtSecret() string {
	return config.Jwt.Secret.String
}

func (r JwtRepository) GenerateRefreshToken() (string, error) {
	return r.GenerateOpaqueToken(64)
}

func (r JwtRepository) GenerateOpaqueToken(size int) (string, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes)
	return token, nil
}
