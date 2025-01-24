package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"crazygames.io/config"
	"crazygames.io/handler/response"

	"crazygames.io/handler/request"
	"crazygames.io/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc services.UserServiceInterface
}

func NewUserHandler(svc services.UserServiceInterface) *UserHandler {
	return &UserHandler{svc: svc}
}

// Create
// @Description Create a new user
// @Tags Users
// @Param UserRequest body request.UserRequest true "Create user request"
// @Accept json
// @Produce json
// @Success 201 {object} entities.User
// @Router /user [post]
func (h *UserHandler) Create(c *gin.Context) {
	var request request.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.svc.Create(&request)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}

// GetAll
// @Description Get all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} []entities.User
// @Router /user [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid page number")
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid limit number")
		return
	}

	users, err := h.svc.GetAll(page, limit)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Calculate total pages
	totalPages := (len(users) + limit - 1) / limit

	response.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", gin.H{
		"totalCount":  len(users),
		"currentPage": page,
		"totalPages":  totalPages,
		"users":       users,
	})
}

// GetByID
// @Description Get user by id
// @Tags Users
// @Param id path uint true "User ID"
// @Accept json
// @Produce json
// @Success 200 {object} entities.User
// @Router /user/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	user, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	response.SuccessResponse(c, http.StatusOK, "successfully", user)
}

// Update
// @Description Update a user
// @Tags Users
// @Param id path uint true "User ID"
// @Param UserUpdateRequest body request.UserUpdateRequest true "Update user request"
// @Accept json
// @Produce json
// @Success 200 {object} entities.User
// @Router /user/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var request request.UserUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.svc.Update(uint(id), &request)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "User updated successfully", user)
}

// Delete
// @Description Update a user
// @Tags Users
// @Param id path uint true "User ID"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Example: {\"message\": \"User deleted successfully\"}"
// @Router /user/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}

// Forgot Password
// @Description Forgot Password
// @Tags Users
// @Param UserEmailRequest body request.UserEmailRequest true "Forgot password Request"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Example: {\"message\": \"The password reset email has been sent.\"}"
// @Router /user/forgot-password [post]
func (h *UserHandler) ForgotPassword(c *gin.Context) {
	var request request.UserEmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.svc.GenerateResetToken(request.Email)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	emailService := services.NewEmailService()

	redirectUrl := config.SMTP.RedirectUrl
	resetLink := redirectUrl + fmt.Sprintf("/reset-password?token=%s", token)

	htmlContent, err := os.ReadFile("templates/email_template.html")
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Parse the template
	tmpl, err := template.New("email").Parse(string(htmlContent))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Replace placeholders
	var body bytes.Buffer
	err = tmpl.Execute(&body, map[string]string{"ResetLink": resetLink})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = emailService.SendEmail(request.Email, "Password Reset", body.String())

	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to send email")
		return
	}

	response.SuccessResponse(c, http.StatusOK, "The password reset email has been sent.", nil)
}

// Reset Password
// @Description Reset Password
// @Tags Users
// @Query token path string true "token"
// @Param UserResetPasswordRequest body request.UserResetPasswordRequest true "Reset password Request"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Example: {\"message\": \"Password reset successful.\"}"
// @Router /user/reset-password [post]
func (h *UserHandler) ResetPassword(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Token is required")
		return
	}
	resetToken, err := h.svc.ValidateResetToken(token)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if resetToken.ExpiredAt.Before(time.Now()) {
		response.ErrorResponse(c, http.StatusBadRequest, "Token has expired")
		return
	}
	if resetToken.IsUsed {
		response.ErrorResponse(c, http.StatusBadRequest, "Token has already been used")
		return
	}

	var request request.UserResetPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.svc.ResetPassword(request.NewPassword, resetToken.Email)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.svc.MarkTokenAsUsed(resetToken.ID)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Password reset successful.", nil)
}
