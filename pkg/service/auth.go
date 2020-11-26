package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
	"time"
)

const (
	salt     = "HHkdsjggsguy"
	signKey  = "JKfhndkkj646yNfsjdjfgfhfJKNnfgh"
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	//передаем ещё на слой ниже в репозиторий
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	//достаем пользователя из БД
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), //валидность токена - 12 часов
			IssuedAt:  time.Now().Unix(),               //время, когда токен был сгенерирован
		}, user.Id})
	return token.SignedString([]byte(signKey))
}
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signKey), nil
	})

	if err != nil {
		return 0, errors.New("wrong token")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("wrong token claims type")
	}
	return claims.UserId, nil
}

func (s *AuthService) CreateCompany(owner models.Company) (int, error) {
	owner.Password = generatePasswordHash(owner.Password)
	//передаем ещё на слой ниже в репозиторий
	return s.repo.CreateCompany(owner)
}
func (s *AuthService) GenerateCompanyOwnerToken(email, password string) (string, error) {
	//достаем пользователя из БД
	user, err := s.repo.GetCompany(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), //валидность токена - 12 часов
			IssuedAt:  time.Now().Unix(),               //время, когда токен был сгенерирован
		}, user.Id})
	return token.SignedString([]byte(signKey))
}
