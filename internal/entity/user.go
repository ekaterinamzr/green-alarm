package entity

const (
	Admin      = 1
	Moderator  = 2
	Authorised = 3
)

type User struct {
	Id         int    `json:"id" db:"id"`
	First_name string `json:"first_name" db:"first_name"`
	Last_name  string `json:"last_name" db:"last_name"`
	Username   string `json:"username" db:"username"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"user_password" db:"user_password"`
	Role       int    `json:"user_role" db:"user_role"`
}
