package dto

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type GetAllUsersResponse struct {
	Users []entity.User
}

type GetUserByIdRequest struct {
	Id int `json:"id"`
}

type GetUserByIdResponse struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"user_password"`
	Role       int    `json:"user_role"`
}

type UpdateUserRequest struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"user_password"`
	Role       int    `json:"user_role"`
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}
