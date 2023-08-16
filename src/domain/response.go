package domain

type PaginateResponse struct {
	Data        interface{} `json:"data"`
	CurrentPage int64       `json:"page"`
	Limit       int64       `json:"limit"`
	TotalPage   int64       `json:"total_pages"`
	TotalData   int64       `json:"total_data"`
}
