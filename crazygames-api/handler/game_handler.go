package handler

import (
	"net/http"
	"strconv"
	"strings"

	"crazygames.io/handler/request"
	"crazygames.io/handler/response"
	"crazygames.io/services"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	svc services.GameServiceInterface
}

func NewGameHandler(svc services.GameServiceInterface) *GameHandler {
	return &GameHandler{svc: svc}
}

// Create
// @Description Create a new game
// @Tags Games
// @Param game_title formData string true "game_title"
// @Param description formData string false "description"
// @Param developer formData string false "developer"
// @Param category_id formData string false "category_id"
// @Param release_date formData string false "release_date"
// @Param thumbnail formData file false "thumbnail"
// @Param technology formData string false "technology"
// @Param rating formData number false "rating"
// @Param hover_video formData file false "hover_video"
// @Param game_url formData string false "game_url"
// @Param play_count formData number false "play_count"
// @Accept multipart/form-data
// @Produce json
// @Success 201 {object} entities.Game
// @Router /game [post]
func (h *GameHandler) Create(c *gin.Context) {
	var request request.GameRequestCreate
	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if request.Thumbnail == nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Thumbnail is required")
		return
	}

	// Validate file type
	if !strings.HasPrefix(request.Thumbnail.Header.Get("Content-Type"), "image/") {
		response.ErrorResponse(c, http.StatusBadRequest, "Thumbnail must be an image")
		return
	}

	if request.HoverVideo == nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Thumbnail is required")
		return
	}

	// Validate file type
	if !strings.HasPrefix(request.HoverVideo.Header.Get("Content-Type"), "video/") {
		response.ErrorResponse(c, http.StatusBadRequest, "Thumbnail must be an video")
		return
	}

	game, err := h.svc.Create(&request)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, game)
}

// GetAll
// @Description Get all games
// @Tags Games
// @Param query query request.GamesRequestQuery true "Query parameters"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=response.GamesResponse}
// @Router /game [get]
func (h *GameHandler) GetAll(c *gin.Context) {
	var query request.GamesRequestQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	games, total, err := h.svc.GetAll(query)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Games retrieved successfully", response.GamesResponse{
		Games:      games,
		Total:      total,
		PageNumber: query.PageNumber,
		PageSize:   query.PageSize,
	})
}

// GetByID
// @Description Get a game by id
// @Tags Games
// @Param id path uint true "Game ID"
// @Accept json
// @Produce json
// @Success 200 {object} entities.Game
// @Router /game/{id} [get]
func (h *GameHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	game, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Game not found")
		return
	}

	c.JSON(http.StatusOK, game)
}

// GetByCategoryID
// @Description Get a game by category id
// @Tags Games
// @Param id path uint true "Category ID"
// @Accept json
// @Produce json
// @Success 200 {object} entities.Game
// @Router /game/category/{id} [get]
func (h *GameHandler) GetByCategoryID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	game, err := h.svc.GetByCategoryID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Game not found")
		return
	}

	c.JSON(http.StatusOK, game)
}

// Update
// @Description Update a game
// @Tags Games
// @Param id path uint true "Game ID"
// @Param game_title formData string true "game_title"
// @Param description formData string false "description"
// @Param developer formData string false "developer"
// @Param category_id formData string false "category_id"
// @Param release_date formData string false "release_date"
// @Param thumbnail formData file false "thumbnail"
// @Param technology formData string false "technology"
// @Param rating formData number false "rating"
// @Param hover_video formData file false "hover_video"
// @Param game_url formData string false "game_url"
// @Param play_count formData number false "play_count"
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} entities.Game
// @Router /game/{id} [put]
func (h *GameHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var request request.GameRequestUpdate
	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validate file type if new thumbnail is provided
	if request.Thumbnail != nil && !strings.HasPrefix(request.Thumbnail.Header.Get("Content-Type"), "image/") {
		response.ErrorResponse(c, http.StatusBadRequest, "Thumbnail must be an image")
		return
	}

	if request.HoverVideo != nil && !strings.HasPrefix(request.HoverVideo.Header.Get("Content-Type"), "video/") {
		response.ErrorResponse(c, http.StatusBadRequest, "Thumbnail must be an video")
		return
	}

	game, err := h.svc.Update(uint(id), &request)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Game updated successfully", game)
}

// Delete
// @Description Delete a game
// @Tags Games
// @Param id path uint true "Game ID"
// @Accept json
// @Produce json
// @Success 200 {object} entities.Game
// @Router /game/{id} [delete]
func (h *GameHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Game deleted successfully", nil)
}

// ListByCategory
// @Description List games by category
// @Tags Games
// @Param category_id path uint true "Category ID"
// @Accept json
// @Produce json
// @Success 200 {array} []entities.Game
// @Router /game/{id} [get]
func (h *GameHandler) ListByCategory(c *gin.Context) {
	categoryId, err := strconv.ParseUint(c.Param("category_id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID")
		return
	}

	games, err := h.svc.ListByCategory(uint(categoryId))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "successfully", games)
}
