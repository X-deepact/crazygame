package request

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role"`
}

type UserUpdateRequest struct {
	Username string `json:"username" binding:"omitempty"`
	Password string `json:"password" binding:"omitempty,min=6"`
	Email    string `json:"email" binding:"omitempty,email"`
	Role     string `json:"role" binding:"omitempty,oneof=player admin"`
}
