package domain

type RequestQuery struct {
	Keyword       string
	Page          int64
	Limit         int64
	SortBy        string
	SortDirection string `default:"DESC"`
	StartDate     string
	EndDate       string
	Filters       map[string]interface{}
}
