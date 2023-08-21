package helpers

import (
	"alqinsidev/jsa-mini-project/aduan/domain"
	"math"
)

func PaginateResponse(query *domain.RequestQuery, data interface{}, totalData int64) *domain.PaginateResponse {
	totalPages := int64(math.Ceil(float64(totalData) / float64(query.Limit)))

	// Return an empty slice [] instead of nil if data is empty
	var returnedData = data
	if d, ok := data.([]domain.AduanTableResponse); ok && len(d) == 0 {
		returnedData = make([]domain.AduanTableResponse, 0)
	}

	return &domain.PaginateResponse{
		Data:        returnedData,
		CurrentPage: query.Page,
		Limit:       query.Limit,
		TotalData:   totalData,
		TotalPage:   totalPages,
	}
}
