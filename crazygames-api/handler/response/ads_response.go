package response

import "crazygames.io/entities"

type AdsResponse struct {
	Ads        []entities.Ads `json:"ads"`
	Total      int64          `json:"total"`
	PageNumber int            `json:"pageNumber"`
	PageSize   int            `json:"pageSize"`
}
