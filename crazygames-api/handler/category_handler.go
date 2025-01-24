package handler

import (
	"crazygames.io/handler/response"
	"net/http"
	"strconv"
	"strings"

	"crazygames.io/entities"
	"crazygames.io/handler/request"
	"crazygames.io/services"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	svc services.CategoryServiceInterface
}

func NewCategoryHandler(svc services.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{
		svc: svc,
	}
}

// Create
// @Description Create a new category
// @Tags Categories
// @Param category_name formData string true "Category Name"
// @Param icon formData file true "Upload Icon"
// @Param description formData string true "Description"
// @Param path formData string true "path"
// @Param is_menu formData string true "is_menu"
// @Accept multipart/form-data
// @Produce json
// @Success 201 {object} entities.Category
// @Router /category [post]
func (h *CategoryHandler) Create(c *gin.Context) {
	var request request.CategoryRequestCreate
	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validate file type
	if !strings.HasPrefix(request.Icon.Header.Get("Content-Type"), "image/") {
		response.ErrorResponse(c, http.StatusBadRequest, "Icon must be an image")
		return
	}
	var category *entities.Category
	var errCreate error
	if category, errCreate = h.svc.CreateCategory(&request); errCreate != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, errCreate.Error())
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetAll
// @Description Get all categories
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {array} []entities.Category
// @Router /category [get]
func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.svc.GetAll()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetMenu
// @Description Get menu
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {array} []entities.Category
// @Router /category/menu [get]
func (h *CategoryHandler) GetMenu(c *gin.Context) {
	categories, err := h.svc.GetMenu()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetByID
// @Description Get category by id
// @Tags Categories
// @Param id path uint true "Category ID"
// @Accept json
// @Produce json
// @Success 200 {object} entities.Category
// @Router /category/{id} [get]
func (h *CategoryHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	category, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Category not found")
		return
	}

	c.JSON(http.StatusOK, category)
}

// Update
// @Description Update a category
// @Tags Categories
// @Param id path uint true "Category ID"
// @Param category_name formData string false "Category Name"
// @Param icon formData file true "Upload Icon"
// @Param description formData string true "Description"
// @Param path formData string true "path"
// @Param is_menu formData string true "is_menu"
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} entities.Category
// @Router /category/{id} [put]
func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var request request.CategoryRequestUpdate
	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Upload new icon if provided
	if request.Icon != nil {
		if !strings.HasPrefix(request.Icon.Header.Get("Content-Type"), "image/") {
			response.ErrorResponse(c, http.StatusBadRequest, "Icon must be an image")
			return
		}
	}

	updatedCategory, err := h.svc.Update(&request, uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Category updated successfully", updatedCategory)
}

// Delete
// @Description Delete a category
// @Tags Categories
// @Param id path uint true "Category ID"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Example: {\"message\": \"Category deleted successfully\"}"
// @Router /category/{id} [delete]
func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Category deleted successfully", nil)
}
