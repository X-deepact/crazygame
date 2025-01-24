package request

// Common request structures that might be shared across handlers

type PaginationRequest struct {
	Page    int `json:"page" form:"page"`
	PerPage int `json:"per_page" form:"per_page"`
}

type IDRequest struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

type SearchRequest struct {
	Query string `json:"query" form:"query"`
}
