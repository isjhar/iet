package repositories

import (
	"isjhar/template/echo-golang/domain/entities"
	"isjhar/template/echo-golang/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

const jwtSecretDefault string = "JWT-SECRET"
const jwtLifeTimeDefault int64 = 60 // minute
const dataKey string = "data"

type JwtRepository struct {
}

func (r JwtRepository) GenerateToken(data interface{}) (string, error) {
	plainToken := jwt.New(jwt.SigningMethodHS512)
	claims := plainToken.Claims.(jwt.MapClaims)
	claims[dataKey] = data
	claims["exp"] = r.getJwtExp()

	securedToken, err := plainToken.SignedString([]byte(r.GetJwtSecret()))
	if err != nil {
		return "", err
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
		return nil, err
	}
	return claims, nil
}

func (r JwtRepository) getJwtExp() int64 {
	jwtlifeTimeString := utils.GetEnvironmentVariable("JWT_LIFE_TIME", strconv.Itoa(int(jwtLifeTimeDefault)))
	jwtLifeTime, err := strconv.Atoi(jwtlifeTimeString)
	if err != nil {
		jwtLifeTime = 0
	}
	return time.Now().Add(time.Minute * time.Duration(jwtLifeTime)).Unix()
}

func (r JwtRepository) GetJwtSecret() string {
	return utils.GetEnvironmentVariable("JWT_SECRET", jwtSecretDefault)
}
