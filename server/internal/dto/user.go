package dto

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"user_password"`
	Role      int    `json:"user_role"`
}

func FromUser(userEntity *entity.User) User {
	return User{
		Id:        userEntity.Id,
		FirstName: userEntity.FirstName,
		LastName:  userEntity.LastName,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Password:  userEntity.Password,
		Role:      userEntity.Role,
	}
}

func FromUsers(userEntities []entity.User) []User {
	dto := make([]User, len(userEntities))
	for i := range dto {
		dto[i] = FromUser(&userEntities[i])
	}
	return dto
}

type GetAllUsersResponse []User

type GetUserByIdRequest struct {
	Id int `json:"id"`
}

type GetUserByIdResponse User

type UpdateUserRequest User

type DeleteUserRequest struct {
	Id int `json:"id"`
}

type ChangeRoleRequest struct {
	Id      int `json:"id"`
	NewRole int `json:"NewRole"`
}
