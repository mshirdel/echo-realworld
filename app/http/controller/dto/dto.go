package dto

type ErrorResponse struct {
	Errors ErrorBody `json:"errors"`
}

type ErrorBody struct {
	Body []string `json:"body"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
