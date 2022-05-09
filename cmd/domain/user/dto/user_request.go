package dto

type UserRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
