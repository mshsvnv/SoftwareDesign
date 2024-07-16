package service

import (
	"context"
	"fmt"
	"src/internal/dto"
	repo "src/internal/repository"
	"src/pkg/logging"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	salt       = "bvelvlerbvlhboge328"
	signingKey = "jkvnljvnlejrnvlebv"
	tokenTTl   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int
}

type IAuthService interface {
	// Register(ctx context.Context, req *dto.RegisterReq)
	GenerateToken(ctx context.Context, req *dto.LoginReq) (string, error)
	ParseToken(token string) (int, error)
}

type AuthService struct {
	logger logging.Interface
	repo   repo.IUserRepository
}

func NewAuthService(logger logging.Interface, repo repo.IUserRepository) *AuthService {
	return &AuthService{
		logger: logger,
		repo:   repo,
	}
}

func (s *AuthService) GenerateToken(ctx context.Context, req *dto.LoginReq) (string, error) {

	s.logger.Infof("login email %s", req.Email)
	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		s.logger.Errorf("get user by email fail, error %s", err.Error())
		return "", fmt.Errorf("get user by email fail, error %s", err)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password))

	if err != nil {
		s.logger.Errorf("wrong password, error %s", err.Error())
		return "", fmt.Errorf("wrong password, error %s", err)
	}

	expiresAt := &jwt.NumericDate{
		time.Now().Add(tokenTTl),
	}

	issuedAt := &jwt.NumericDate{
		time.Now(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, err := token.Method.(*jwt.SigningMethodHMAC)
		if !err {
			s.logger.Errorf("invalid signing method, error %s", err)
			return -1, fmt.Errorf("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		s.logger.Errorf("%s", err)
		return -2, err
	}

	claims := token.Claims.(*tokenClaims)
	if claims == nil {
		s.logger.Errorf("token claims are not of type *tokenClaims")
		return -3, fmt.Errorf("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}
