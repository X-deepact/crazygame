package handler

import (
	"net/http"

	"crazygames.io/handler/request"
	"crazygames.io/handler/response"
	"crazygames.io/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authSvc services.AuthServiceInterface
	userSvc services.UserServiceInterface
}

func NewAuthHandler(authSvc services.AuthServiceInterface, userSvc services.UserServiceInterface) *AuthHandler {
	return &AuthHandler{authSvc: authSvc, userSvc: userSvc}
}

// Login
// @Description Log in an existing user
// @Tags Auth
// @Param LoginRequest body request.LoginRequest true "Login request"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Example: {\"message\": \"Login successful\", \"token\": \"your-jwt-token\"}"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest request.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	token, err := h.authSvc.Login(&loginRequest)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	response.SuccessResponse(c, http.StatusOK, "email_exists", token)

}

// Register
// @Description Register a new user
// @Tags Auth
// @Param RegisterRequest body request.RegisterRequest true "Register request"
// @Accept json
// @Produce json
// @Success 201 {object} entities.User
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var registerRequest request.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authSvc.Register(&registerRequest)
	if err != nil {
		response.ErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}

// CheckEmail
// @Description Check if email exists and respond accordingly
// @Tags Authentication
// @Param UserEmailRequest body request.UserEmailRequest true "Email request"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Example: {\"message\": \"email_exists\"}"
// @Router /auth/check-email [post]
func (h *AuthHandler) CheckEmail(c *gin.Context) {
	var request request.UserEmailRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	user, err := h.userSvc.GetByEmail(request.Email)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "email_not_found")
		return
	}

	response.SuccessResponse(c, http.StatusOK, "email_exists", user)
}
