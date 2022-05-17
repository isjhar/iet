package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const jwtSecretDefault string = "JWT-SECRET"
const jwtLifeTimeDefault int64 = 30

func GenerateJWT(user interface{}) (string, error) {
	plainToken := jwt.New(jwt.SigningMethodHS512)
	claims := plainToken.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = GetJwtExp()

	securedToken, err := plainToken.SignedString([]byte(GetJwtSecret()))
	if err != nil {
		return "", err
	}

	return securedToken, nil

}

func GetJwtExp() int64 {
	jwtlifeTimeString := GetEnvironmentVariable("JWT_LIFE_TIME", strconv.Itoa(int(jwtLifeTimeDefault)))
	jwtLifeTime, err := strconv.Atoi(jwtlifeTimeString)
	if err != nil {
		jwtLifeTime = 0
	}
	return time.Now().Add(time.Minute * time.Duration(jwtLifeTime)).Unix()
}

func GetJwtSecret() string {
	return GetEnvironmentVariable("JWT_SECRET", jwtSecretDefault)
}
