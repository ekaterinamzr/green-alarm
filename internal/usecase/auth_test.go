package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
	mock "github.com/ekaterinamzr/green-alarm/internal/infrastructure/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var errAuthInternal = errors.New("internal server error")

const passwordHash = "PasswordHash"
const token = "Token"

func hash(password, salt string) string {
	return passwordHash
}

func TestAuthSignUp(t *testing.T) {
	type args struct {
		ctx     context.Context
		request dto.SignUpRequest
	}

	tests := []struct {
		name  string
		args  args
		repo  func(ctrl *gomock.Controller) *mock.MockAuthRepository
		token func(ctrl *gomock.Controller) *mock.MockTokenizer
		res   *dto.SignUpResponse
		err   error
	}{
		{
			name: "success",
			args: args{ctx: context.Background(), request: dto.SignUpRequest{Username: "testUser", Password: "testPassword"}},
			repo: func(ctrl *gomock.Controller) *mock.MockAuthRepository {
				m := mock.NewMockAuthRepository(ctrl)
				m.EXPECT().Create(gomock.Any(), entity.User{Username: "testUser", Password: passwordHash, Role: entity.Authorised}).Return(1, nil).Times(1)
				return m
			},
			token: func(ctrl *gomock.Controller) *mock.MockTokenizer {
				return mock.NewMockTokenizer(ctrl)
			},
			res: &dto.SignUpResponse{Id: 1},
			err: nil,
		},
		{
			name: "repo error",
			args: args{ctx: context.Background(), request: dto.SignUpRequest{Username: "testUser", Password: "testPassword"}},
			repo: func(ctrl *gomock.Controller) *mock.MockAuthRepository {
				m := mock.NewMockAuthRepository(ctrl)
				m.EXPECT().Create(gomock.Any(), entity.User{Username: "testUser", Password: passwordHash, Role: entity.Authorised}).Return(0, errAuthInternal).Times(1)
				return m
			},
			token: func(ctrl *gomock.Controller) *mock.MockTokenizer {
				return mock.NewMockTokenizer(ctrl)
			},
			res: nil,
			err: errAuthInternal,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := NewAuthUseCase(tc.repo(ctrl), tc.token(ctrl), hash, "salt")

			res, err := uc.SignUp(tc.args.ctx, tc.args.request)

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestAuthSignIn(t *testing.T) {
	type args struct {
		ctx     context.Context
		request dto.SignInRequest
	}

	tests := []struct {
		name  string
		args  args
		repo  func(ctrl *gomock.Controller) *mock.MockAuthRepository
		token func(ctrl *gomock.Controller) *mock.MockTokenizer
		res   *dto.SignInResponse
		err   error
	}{
		{
			name: "success",
			args: args{ctx: context.Background(), request: dto.SignInRequest{Username: "testUser", Password: "testPassword"}},
			repo: func(ctrl *gomock.Controller) *mock.MockAuthRepository {
				m := mock.NewMockAuthRepository(ctrl)
				m.EXPECT().GetUser(gomock.Any(), "testUser", passwordHash).Return(&entity.User{Id: 1, Role: 1}, nil).Times(1)
				return m
			},
			token: func(ctrl *gomock.Controller) *mock.MockTokenizer {
				m := mock.NewMockTokenizer(ctrl)
				m.EXPECT().GenerateToken(gomock.Any(), 1, 1).Return(token, nil)
				return m
			},
			res: &dto.SignInResponse{Id: 1, Role: 1, Token: token},
			err: nil,
		},
		{
			name: "repo error",
			args: args{ctx: context.Background(), request: dto.SignInRequest{Username: "testUser", Password: "testPassword"}},
			repo: func(ctrl *gomock.Controller) *mock.MockAuthRepository {
				m := mock.NewMockAuthRepository(ctrl)
				m.EXPECT().GetUser(gomock.Any(), "testUser", passwordHash).Return(nil, errAuthInternal).Times(1)
				return m
			},
			token: func(ctrl *gomock.Controller) *mock.MockTokenizer {
				m := mock.NewMockTokenizer(ctrl)
				return m
			},
			res: nil,
			err: errAuthInternal,
		},
		{
			name: "token error",
			args: args{ctx: context.Background(), request: dto.SignInRequest{Username: "testUser", Password: "testPassword"}},
			repo: func(ctrl *gomock.Controller) *mock.MockAuthRepository {
				m := mock.NewMockAuthRepository(ctrl)
				m.EXPECT().GetUser(gomock.Any(), "testUser", passwordHash).Return(&entity.User{Id: 1, Role: 1}, nil).Times(1)
				return m
			},
			token: func(ctrl *gomock.Controller) *mock.MockTokenizer {
				m := mock.NewMockTokenizer(ctrl)
				m.EXPECT().GenerateToken(gomock.Any(), 1, 1).Return("", errAuthInternal)
				return m
			},
			res: nil,
			err: errAuthInternal,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := NewAuthUseCase(tc.repo(ctrl), tc.token(ctrl), hash, "salt")

			res, err := uc.SignIn(tc.args.ctx, tc.args.request)

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
