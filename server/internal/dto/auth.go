package dto

type SignUpRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"user_password"`
}

type SignUpResponse struct {
	Id int `json:"id"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"user_password"`
}

type SignInResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
	Role  int    `json:"role"`
}