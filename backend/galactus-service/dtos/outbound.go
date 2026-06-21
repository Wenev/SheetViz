package dtos

type UserResponse struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

type UserListResponse struct {
	Data []UserResponse `json:"data"`
}