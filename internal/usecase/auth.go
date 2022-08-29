package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

// Auth use case interface
type Auth interface {
	SignUp(context.Context, dto.SignUpRequest) (int, error)
}

type AuthRepository interface {
	CreateUser(context.Context, entity.User) (int, error)
}

type AuthUseCase struct {
	repo AuthRepository
	salt string
}

func NewAuthUseCase(r AuthRepository, s string) *AuthUseCase {
	return &AuthUseCase{repo: r, salt: s}
}

func (uc *AuthUseCase) SignUp(ctx context.Context, u dto.SignUpRequest) (int, error) {
	u.Password = uc.generatePasswordHash(u.Password)

	id, err := uc.repo.CreateUser(ctx, entity.User{
		First_name: u.First_name,
		Last_name: u.Last_name,
		Username: u.Username,
		Email: u.Email,
		Password: u.Password,
		Role: entity.Authorised,
	})
	
	if err != nil {
		return 0, fmt.Errorf("AuthUseCase - SignUp - uc.repo.CreateUser: %w", err)
	}

	return id, nil
}

func (uc *AuthUseCase) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(uc.salt)))
}
