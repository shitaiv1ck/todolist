package users_transport

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=8,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type CreateUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type GetUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
