package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type AuthRepository interface {
	Create(context.Context, entity.User) (string, error)
	GetUser(context.Context, string, string) (*entity.User, error)
}

type TokenService interface {
	GenerateToken(ctx context.Context, id string, role int) (string, error)
	ParseToken(context.Context, string) (id string, role int, err error)
}

type AuthUseCase struct {
	repo  AuthRepository
	token TokenService
	salt  string
}

func NewAuthUseCase(r AuthRepository, t TokenService, salt string) *AuthUseCase {
	return &AuthUseCase{
		repo:  r,
		salt:  salt,
		token: t,
	}
}

func (uc *AuthUseCase) SignUp(ctx context.Context, u dto.SignUpRequest) (*dto.SignUpResponse, error) {
	u.Password = uc.generatePasswordHash(u.Password)

	id, err := uc.repo.Create(ctx, entity.User{
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

// TODO: move to infrastructure
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

	token, err := uc.token.GenerateToken(ctx, user.Id, user.Role)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - SignIn - token.SignedString: %w", err)
	}

	return &dto.SignInResponse{Id: user.Id, Role: user.Role, Token: token}, err
}

func (uc *AuthUseCase) ParseToken(ctx context.Context, token string) (string, int, error) {
	id, role, err := uc.token.ParseToken(ctx, token)
	return id, role, err
}
