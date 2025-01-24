package handler

import (
	"fmt"
	"net/http"

	"crazygames.io/config"
	"crazygames.io/handler/response"
	"crazygames.io/services"
	"crazygames.io/utils"
	"github.com/gin-gonic/gin"
)

type OAuthHandler struct {
	svc services.OAuthServiceInterface
}

func NewOAuthHandler(svc services.OAuthServiceInterface) *OAuthHandler {
	return &OAuthHandler{
		svc: svc,
	}
}

// Google Login Handler
func (h *OAuthHandler) GoogleLogin(c *gin.Context) {
	authURL := h.svc.GetGoogleAuthURL("state-token")
	c.JSON(http.StatusOK, gin.H{"redirect_url": authURL})
}

// Google Callback Handler
func (h *OAuthHandler) GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "state-token" {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid state")
		return
	}

	code := c.Query("code")
	if code == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Missing code parameter")
		return
	}

	token, err := h.svc.ExchangeCodeForToken(c.Request.Context(), code)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Fetch user info
	userInfo, err := utils.FetchUserInfo(token.AccessToken)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Save user to database
	if err := h.svc.HandleGoogleUser(userInfo); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	queryStrings := fmt.Sprintf("?token=%v&uid=%v&uname=%v", token.AccessToken, userInfo["email"], userInfo["name"])
	c.Redirect(http.StatusTemporaryRedirect, config.AppConfig.FRONTEND_URL+queryStrings)
}
