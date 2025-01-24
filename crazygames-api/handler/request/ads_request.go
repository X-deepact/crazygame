package request

import "mime/multipart"

type AdsRequestCreate struct {
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	Position uint                  `form:"position" binding:"required"`
	GameId   uint                  `form:"game_id" binding:"required"`
}

type AdsRequestUpdate struct {
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	Position uint                  `form:"position" binding:"required"`
	GameId   uint                  `form:"game_id" binding:"required"`
}

type AdsRequestQuery struct {
	PageNumber int `form:"page_number" binding:"required,min=1"`
	PageSize   int `form:"page_size" binding:"required,min=1,max=100"`
}
