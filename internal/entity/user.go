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
	Id        int    `json:"id" db:"id" bson:"_id"`
	FirstName string `json:"first_name" db:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" db:"last_name" bson:"last_name"`
	Username  string `json:"username" db:"username" bson:"username"`
	Email     string `json:"email" db:"email" bson:"email"`
	Password  string `json:"user_password" db:"user_password" bson:"user_password"`
	Role      int    `json:"user_role" db:"user_role" bson:"user_role"`
}
