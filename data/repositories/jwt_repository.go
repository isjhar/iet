package repositories

import (
	"isjhar/template/echo-golang/domain/entities"
	"isjhar/template/echo-golang/utils"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"gopkg.in/guregu/null.v4"
)

const jwtSecretDefault string = "JWT-SECRET"
const dataKey string = "data"

type JwtRepository struct {
}

func (r JwtRepository) GenerateToken(data interface{}, expMinutes null.Int) (string, error) {
	plainToken := jwt.New(jwt.SigningMethodHS512)
	claims := plainToken.Claims.(jwt.MapClaims)
	claims[dataKey] = data
	if expMinutes.Valid {
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(expMinutes.Int64)).Unix()
	}

	securedToken, err := plainToken.SignedString([]byte(r.GetJwtSecret()))
	if err != nil {
		log.Printf("error sign token: %v", err)
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
		log.Printf("error validate token: %v", err)
		return nil, entities.InternalServerError
	}
	return claims, nil
}

func (r JwtRepository) GetJwtSecret() string {
	return utils.GetEnvironmentVariable("JWT_SECRET", jwtSecretDefault)
}
