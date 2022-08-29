package dto

type SignUpRequest struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"user_password"`
}

type SignUpResponse struct {
	Id int `json:"id"`
}
