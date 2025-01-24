package response

import "crazygames.io/entities"

type GamesResponse struct {
	Games      []entities.Game `json:"games"`
	Total      int64           `json:"total"`
	PageNumber int             `json:"pageNumber"`
	PageSize   int             `json:"pageSize"`
}
