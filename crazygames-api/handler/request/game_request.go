package request

import "mime/multipart"

type GameRequestCreate struct {
	GameTitle   string                `form:"game_title" binding:"required"`
	Description string                `form:"description"`
	Developer   string                `form:"developer"`
	CategoryID  string                `form:"category_id"`
	ReleaseDate string                `form:"release_date"`
	Thumbnail   *multipart.FileHeader `form:"thumbnail"`
	Technology  string                `form:"technology"`
	Rating      float64               `form:"rating"`
	HoverVideo  *multipart.FileHeader `form:"hover_video"`
	GameURL     string                `form:"game_url"`
	PlayCount   int                   `form:"play_count"`
}

type GameRequestUpdate struct {
	GameTitle   string                `form:"game_title" binding:"required"`
	Description string                `form:"description"`
	Developer   string                `form:"developer"`
	CategoryID  string                `form:"category_id"`
	ReleaseDate string                `form:"release_date"`
	Thumbnail   *multipart.FileHeader `form:"thumbnail"`
	Technology  string                `form:"technology"`
	Rating      float64               `form:"rating"`
	HoverVideo  *multipart.FileHeader `form:"hover_video"`
	GameURL     string                `form:"game_url"`
	PlayCount   int                   `form:"play_count"`
}

type GamesRequestQuery struct {
	PageNumber int    `form:"page_number" binding:"required,min=1"`
	PageSize   int    `form:"page_size" binding:"required,min=1,max=100"`
	Search     string `form:"search"`
}
