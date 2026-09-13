package dto

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	Message string `json:message`
	Name    string `json:name`
}

type UserResponse struct {
	ID   int    `json:id`
	Name string `json:name`
}

type ErrorResponse struct {
	Error string `json:error`
}

type User struct {
	ID   int
	Name string
}
