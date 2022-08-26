package repository_pg

import "github.com/ekaterinamzr/green-alarm/internal/entity"

type UserRepository struct {

}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetById(id int) (*entity.User, error) {
	return &entity.User{}, nil
}