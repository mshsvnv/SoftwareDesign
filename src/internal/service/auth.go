package service

import (
	"context"
	"fmt"
	"src/internal/dto"
	"src/internal/model"
	repo "src/internal/repository"
	"src/pkg/logging"
	"src/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int
}

type IAuthService interface {
	Login(ctx context.Context, req *dto.LoginReq) (string, error)
	Register(ctx context.Context, req *dto.RegisterReq) (string, error)
	GenerateToken(userID int) (string, error)
	ParseToken(token string) (int, error)
}

type AuthService struct {
	logger logging.Interface
	repo   repo.IUserRepository

	signingKey     string
	accessTokenTTL time.Duration
}

func NewAuthService(
	logger logging.Interface,
	repo repo.IUserRepository,
	signingKey string,
	accessTokenTTL time.Duration,
) *AuthService {
	return &AuthService{
		logger:         logger,
		repo:           repo,
		signingKey:     signingKey,
		accessTokenTTL: accessTokenTTL,
	}
}

func (s *AuthService) Login(ctx context.Context, req *dto.LoginReq) (string, error) {

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

	return s.GenerateToken(user.ID)
}

func (s *AuthService) Register(ctx context.Context, req *dto.RegisterReq) (string, error) {

	s.logger.Infof("register email %s", req.Email)
	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if user != nil {
		s.logger.Errorf("get user by email fail, error %s", err.Error())
		return "", fmt.Errorf("get user by email fail, error %s", err)
	}

	user = &model.User{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Role:     model.UserRoleCustomer,
		Password: utils.HashAndSalt([]byte(req.Password)),
	}

	err = s.repo.Create(ctx, user)

	if err != nil {
		s.logger.Errorf("create fail, error %s", err.Error())
		return "", fmt.Errorf("create fail, error %s", err)
	}

	return s.GenerateToken(user.ID)
}

func (s *AuthService) GenerateToken(userID int) (string, error) {
	expiresAt := &jwt.NumericDate{
		time.Now().Add(s.accessTokenTTL),
	}

	issuedAt := &jwt.NumericDate{
		time.Now(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
		userID,
	})

	return token.SignedString([]byte(s.signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, err := token.Method.(*jwt.SigningMethodHMAC)
		if !err {
			s.logger.Errorf("invalid signing method, error %s", err)
			return 0, fmt.Errorf("invalid signing method")
		}

		return []byte(s.signingKey), nil
	})

	if err != nil {
		s.logger.Errorf("%s", err)
		return 0, err
	}

	claims := token.Claims.(*tokenClaims)
	if claims == nil {
		s.logger.Errorf("token claims are not of type *tokenClaims")
		return 0, fmt.Errorf("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}
