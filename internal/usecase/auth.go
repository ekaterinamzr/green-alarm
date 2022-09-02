package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

// Auth use case interface
type Auth interface {
	SignUp(context.Context, dto.SignUpRequest) (*dto.SignUpResponse, error)
	SignIn(context.Context, dto.SignInRequest) (*dto.SignInResponse, error)
}

type AuthRepository interface {
	CreateUser(context.Context, entity.User) (int, error)
	GetUser(context.Context, string, string) (*entity.User, error)
}

type AuthUseCase struct {
	repo       AuthRepository
	salt       string
	tokenTTL   time.Duration
	signingKey string
}

type UserJWTClaims struct {
	jwt.StandardClaims
	UserId int  `json:"user_id"`
	UserRole int `json:"user_role"`
}

func NewAuthUseCase(r AuthRepository, salt, signingKey string, tokenTTL int) *AuthUseCase {
	return &AuthUseCase{
		repo:       r,
		salt:       salt,
		signingKey: signingKey,
		tokenTTL:   time.Duration(tokenTTL),
	}
}

func (uc *AuthUseCase) SignUp(ctx context.Context, u dto.SignUpRequest) (*dto.SignUpResponse, error) {
	u.Password = uc.generatePasswordHash(u.Password)

	id, err := uc.repo.CreateUser(ctx, entity.User{
		First_name: u.First_name,
		Last_name:  u.Last_name,
		Username:   u.Username,
		Email:      u.Email,
		Password:   u.Password,
		Role:       entity.Authorised,
	})

	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - SignUp - uc.repo.CreateUser: %w", err)
	}

	return &dto.SignUpResponse{Id: id}, nil
}

func (uc *AuthUseCase) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(uc.salt)))
}

func (uc *AuthUseCase) SignIn(ctx context.Context, u dto.SignInRequest) (*dto.SignInResponse, error) {
	user, err := uc.repo.GetUser(ctx, u.Username, uc.generatePasswordHash(u.Password))
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - SignIn - uc.repo.GetUser: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserJWTClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(uc.tokenTTL * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Role,
	})

	signedToken, err := token.SignedString([]byte(uc.signingKey))
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - SignIn - token.SignedString: %w", err)
	}

	return &dto.SignInResponse{Id: user.Id, Role: user.Role, Token: signedToken}, err
}
