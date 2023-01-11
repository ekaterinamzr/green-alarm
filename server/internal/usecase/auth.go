package usecase

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type hashFync func(password, salt string) string

type AuthUseCase struct {
	repo  AuthRepository
	token Tokenizer
	hash  hashFync
	salt  string
}

func NewAuthUseCase(r AuthRepository, t Tokenizer, h hashFync, salt string) *AuthUseCase {
	return &AuthUseCase{
		repo:  r,
		salt:  salt,
		token: t,
		hash:  h,
	}
}

func (uc *AuthUseCase) SignUp(ctx context.Context, u dto.SignUpRequest) (*dto.SignUpResponse, error) {
	u.Password = uc.hash(u.Password, uc.salt)

	id, err := uc.repo.Create(ctx, entity.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Role:      entity.Authorised,
	})

	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - SignUp - uc.repo.CreateUser: %w", err)
	}

	return &dto.SignUpResponse{Id: id}, nil
}

func (uc *AuthUseCase) SignIn(ctx context.Context, u dto.SignInRequest) (*dto.SignInResponse, error) {
	user, err := uc.repo.GetUser(ctx, u.Username, uc.hash(u.Password, uc.salt))
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - SignIn - uc.repo.GetUser: %w", err)
	}

	token, err := uc.token.GenerateToken(ctx, user.Id, user.Role)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - SignIn - token.SignedString: %w", err)
	}

	return &dto.SignInResponse{Id: user.Id, Role: user.Role, Token: token}, err
}

func (uc *AuthUseCase) ParseToken(ctx context.Context, token string) (id int, role int, err error) {
	return uc.token.ParseToken(ctx, token)
}
