package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"crazygames.io/entities"
	"crazygames.io/handler/request"
	"crazygames.io/handler/response"
	"crazygames.io/services"
)

type AdsHandler struct {
	svc services.AdsServiceInterface
}

func NewAdsHandler(svc services.AdsServiceInterface) *AdsHandler {
	return &AdsHandler{
		svc: svc,
	}
}

// Create
// @Description
// @Tags Advertisements
// @Param image formData file true "Upload Image"
// @Param position formData uint true "Position"
// @Param game_id formData uint true "Game ID"
// @Accept multipart/form-data
// @Produce json
// @Success 201 {object} entities.Ads
// @Router /ads [post]
func (h *AdsHandler) Create(c *gin.Context) {
	var request request.AdsRequestCreate
	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validate file type
	if !strings.HasPrefix(request.Image.Header.Get("Content-Type"), "image/") {
		response.ErrorResponse(c, http.StatusBadRequest, "Image must be an image")
		return
	}
	var ads *entities.Ads
	var errCreate error
	if ads, errCreate = h.svc.Create(&request); errCreate != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, errCreate.Error())
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "Advertisement created successfully", ads)
}

// GetAll
// @Description Get all advertisements
// @Tags Advertisements
// @Param query query request.AdsRequestQuery true "Query parameters"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=response.AdsResponse} "Advertisements retrieved successfully"
// @Router /ads [get]
func (h *AdsHandler) GetAll(c *gin.Context) {
	var query request.AdsRequestQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ads, total, err := h.svc.GetAll(query)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Advertisements retrieved successfully", response.AdsResponse{
		Ads:        ads,
		Total:      total,
		PageNumber: query.PageNumber,
		PageSize:   query.PageSize,
	})
}

// GetByID
// @Description Get advertisement by id
// @Tags Advertisements
// @Param id path uint true "Ads ID"
// @Accept json
// @Produce json
// @Success 200 {object} entities.Ads
// @Router /ads/{id} [get]
func (h *AdsHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	ads, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Advertisement not found")
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Advertisement retrieved successfully", ads)
}

// Update
// @Description Update advertisement by id
// @Tags Advertisements
// @Param id path uint true "Ads ID"
// @Param image formData file true "Upload Image"
// @Param position formData uint true "Position"
// @Param game_id formData uint true "Game ID"
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} entities.Ads
// @Router /ads/{id} [put]
func (h *AdsHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var request request.AdsRequestUpdate
	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Upload new image if provided
	if request.Image != nil {
		if !strings.HasPrefix(request.Image.Header.Get("Content-Type"), "image/") {
			response.ErrorResponse(c, http.StatusBadRequest, "Image must be an image")
			return
		}
	}

	updatedAds, err := h.svc.Update(&request, uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Advertisement updated successfully", updatedAds)
}

// Delete
// @Description Delete advertisement by id
// @Tags Advertisements
// @Param id path uint true "Ads ID"
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} map[string]string "Example: {\"message\": \"success\"}"
// @Router /ads/{id} [delete]
func (h *AdsHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Advertisement deleted successfully", nil)
}
