package request

type UserEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UserResetPasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
