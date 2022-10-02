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

var errIncidentInternal = errors.New("internal server error")

func TestIncidentGetById(t *testing.T) {
	type args struct {
		ctx     context.Context
		request dto.GetIncidentByIdRequest
	}

	tests := []struct {
		name string
		repo func(ctrl *gomock.Controller) *mock.MockIncidentRepository
		args args
		res  *dto.GetIncidentByIdResponse
		err  error
	}{
		{
			name: "success",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().GetById(gomock.Any(), 1).Return(&entity.Incident{Id: 1}, nil).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.GetIncidentByIdRequest{Id: 1}},
			res:  &dto.GetIncidentByIdResponse{Id: 1},
			err:  nil,
		},
		{
			name: "no such id",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().GetById(gomock.Any(), -2).Return(nil, errIncidentInternal).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.GetIncidentByIdRequest{Id: -2}},
			res:  nil,
			err:  errIncidentInternal,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := NewIncidentUseCase(tc.repo(ctrl))

			res, err := uc.GetById(tc.args.ctx, tc.args.request)

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestIncidentGetByType(t *testing.T) {
	type args struct {
		ctx     context.Context
		request dto.GetIncidentsByTypeRequest
	}

	tests := []struct {
		name string
		repo func(ctrl *gomock.Controller) *mock.MockIncidentRepository
		args args
		res  *dto.GetIncidentsResponse
		err  error
	}{
		{
			name: "success",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().GetByType(gomock.Any(), 1).Return([]entity.Incident{{Type: 1}, {Type: 1}}, nil).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.GetIncidentsByTypeRequest{IncidentType: 1}},
			res:  &dto.GetIncidentsResponse{{Type: 1}, {Type: 1}},
			err:  nil,
		},
		{
			name: "no such type",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().GetByType(gomock.Any(), -2).Return(nil, errIncidentInternal).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.GetIncidentsByTypeRequest{IncidentType: -2}},
			res:  nil,
			err:  errIncidentInternal,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := NewIncidentUseCase(tc.repo(ctrl))

			res, err := uc.GetByType(tc.args.ctx, tc.args.request)

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestIncidentGetAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name string
		repo func(ctrl *gomock.Controller) *mock.MockIncidentRepository
		args args
		res  *dto.GetIncidentsResponse
		err  error
	}{
		{
			name: "success",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().GetAll(gomock.Any()).Return([]entity.Incident{}, nil).Times(1)
				return m
			},
			args: args{ctx: context.Background()},
			res:  &dto.GetIncidentsResponse{},
			err:  nil,
		},
		{
			name: "repo error",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().GetAll(gomock.Any()).Return(nil, errIncidentInternal).Times(1)
				return m
			},
			args: args{ctx: context.Background()},
			res:  nil,
			err:  errIncidentInternal,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := NewIncidentUseCase(tc.repo(ctrl))

			res, err := uc.GetAll(tc.args.ctx)

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestIncidentCreate(t *testing.T) {
	type args struct {
		ctx     context.Context
		request dto.CreateIncidentRequest
	}

	tests := []struct {
		name string
		repo func(ctrl *gomock.Controller) *mock.MockIncidentRepository
		args args
		res  *dto.CreateIncidentResponse
		err  error
	}{
		{
			name: "success",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(1, nil).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.CreateIncidentRequest{}},
			res:  &dto.CreateIncidentResponse{Id: 1},
			err:  nil,
		},
		{
			name: "repo error",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(0, errIncidentInternal).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.CreateIncidentRequest{}},
			res:  nil,
			err:  errIncidentInternal,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := NewIncidentUseCase(tc.repo(ctrl))

			res, err := uc.Create(tc.args.ctx, tc.args.request)

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestIncidentUpdate(t *testing.T) {
	type args struct {
		ctx     context.Context
		request dto.UpdateIncidentRequest
	}

	tests := []struct {
		name string
		repo func(ctrl *gomock.Controller) *mock.MockIncidentRepository
		args args
		err  error
	}{
		{
			name: "success",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().Update(gomock.Any(), entity.Incident{Id: 1}).Return(nil).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.UpdateIncidentRequest{Id: 1}},
			err:  nil,
		},
		{
			name: "repo error",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().Update(gomock.Any(), entity.Incident{Id: 1, Longitude: 250}).Return(errIncidentInternal).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.UpdateIncidentRequest{Id: 1, Longitude: 250}},
			err:  errIncidentInternal,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := NewIncidentUseCase(tc.repo(ctrl))

			err := uc.Update(tc.args.ctx, tc.args.request)

			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestIncidentDelete(t *testing.T) {
	type args struct {
		ctx     context.Context
		request dto.DeleteIncidentRequest
	}

	tests := []struct {
		name string
		repo func(ctrl *gomock.Controller) *mock.MockIncidentRepository
		args args
		err  error
	}{
		{
			name: "success",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().Delete(gomock.Any(), 1).Return(nil).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.DeleteIncidentRequest{Id: 1}},
			err:  nil,
		},
		{
			name: "repo error",
			repo: func(ctrl *gomock.Controller) *mock.MockIncidentRepository {
				m := mock.NewMockIncidentRepository(ctrl)
				m.EXPECT().Delete(gomock.Any(), 1).Return(errIncidentInternal).Times(1)
				return m
			},
			args: args{ctx: context.Background(), request: dto.DeleteIncidentRequest{Id: 1}},
			err:  errIncidentInternal,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := NewIncidentUseCase(tc.repo(ctrl))

			err := uc.Delete(tc.args.ctx, tc.args.request)

			require.ErrorIs(t, err, tc.err)
		})
	}
}
