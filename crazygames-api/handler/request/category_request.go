package request

import "mime/multipart"

type CategoryRequestCreate struct {
	CategoryName string                `form:"category_name" binding:"required"`
	Description  string                `form:"description" binding:"required"`
	Icon         *multipart.FileHeader `form:"icon" binding:"required"`
	Path         string                `form:"path" binding:"required"`
	IsMenu       string                `form:"is_menu" binding:"required"`
}

type CategoryRequestUpdate struct {
	CategoryName string                `form:"category_name"`
	Description  string                `form:"description" binding:"required"`
	Icon         *multipart.FileHeader `form:"icon"`
	Path         string                `form:"path" binding:"required"`
	IsMenu       string                `form:"is_menu" binding:"required"`
}
