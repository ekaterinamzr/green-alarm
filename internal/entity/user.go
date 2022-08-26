package entity 

type User struct {
	id int `json:"id"`
	first_name string `json:"first_name"`
	last_name string `json:"last_name"`
	username string `json:"username"`
	email string `json:"email"`
	password string `json:"user_password"`
	role string `json:"user_role"`
}