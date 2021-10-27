package service

import (
	"EStore/models"
	"EStore/pkg/repository"
	"EStore/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

// Структура для создания токена
type tokenClaims struct {
	jwt.StandardClaims
	UserId     int `json:"user_id"`
	UserTypeId int `json:"user_type"`
}

// Генерация токена
func GenerateToken(username, password string) (string, error) {
	user, err := repository.GetUser(username, utils.GenerateSha1(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Server",
		},
		user.Id,
		user.TypeId,
	})

	return token.SignedString([]byte(signingKey))
}

// Рассшифровка токена
func ParseToken(accessToken string) (user models.User, err error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return models.User{}, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return models.User{}, errors.New("token claims are not type of *tokenClaims")
	}

	user.Id = claims.UserId
	user.TypeId = claims.UserTypeId
	return user, nil
}

// Создание пользователя
func CreateUser(user models.User) (int, error) {
	user.Password = utils.GenerateSha1(user.Password)

	typeId, exists := repository.TypeExists("user")
	if !exists {
		user.TypeId = 1
	} else {
		user.TypeId = typeId
	}

	userId, err := repository.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
