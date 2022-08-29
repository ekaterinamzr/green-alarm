package entity

const (
	Admin      = 1
	Moderator  = 2
	Authorised = 3
)

type User struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"user_password"`
	Role       int    `json:"user_role"`
}
