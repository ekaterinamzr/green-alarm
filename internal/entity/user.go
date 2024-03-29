package entity

const (
	Admin = iota + 1
	Moderator
	Authorised
)

type UserRole struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"role_name" db:"role_name"`
}

type User struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"user_password" db:"user_password"`
	Role      int    `json:"user_role" db:"user_role"`
}
