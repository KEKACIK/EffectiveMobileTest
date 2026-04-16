package api

type PaginationResponse struct {
	Total int   `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Items []any `json:"items"`
}
