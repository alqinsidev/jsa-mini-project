package domain

import (
	"time"

	"github.com/google/uuid"
)

type Aduan struct {
	ID                     uuid.UUID `json:"id"`
	ComplainID             string    `json:"complain_id"`
	Title                  string    `json:"title"`
	Description            string    `json:"description"`
	Evidence               string    `json:"evidence"`
	ComplainantID          uuid.UUID `json:"complainant_id"`
	ComplainantPositionLat string    `json:"complainant_position_lat"`
	ComplainantPositionLon string    `json:"complainant_position_lon"`
	ComplainedCity         string    `json:"complained_city"`
	ComplainedDistrict     string    `json:"complained_district"`
	ComplainedSubDistrict  string    `json:"complained_sub_district"`
	ComplainedAddress      string    `json:"complained_address"`
	ComplainedGMapLink     string    `json:"complained_gmap_link"`
	CategoryID             uuid.UUID `json:"category_id"`
	SubCategoryID          uuid.UUID `json:"sub_category_id"`
	Status                 int       `json:"status"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type AduanDetail struct {
	ID                         uuid.UUID `json:"id"`
	ComplainID                 string    `json:"complain_id"`
	Title                      string    `json:"title"`
	Description                string    `json:"description"`
	Evidence                   string    `json:"evidence"`
	ComplainantID              uuid.UUID `json:"complainant_id"`
	ComplainantName            string    `json:"complainant_name"`
	ComplainantPhone           string    `json:"complainant_phone"`
	ComplainantSocialMediaID   uuid.UUID `json:"complainant_social_media_id"`
	ComplainantSocialMediaName string    `json:"complainant_social_media_name"`
	ComplainantSocialMediaLink string    `json:"complainant_social_media_link"`
	ComplainantPositionLat     float64   `json:"complainant_position_lat"`
	ComplainantPositionLon     float64   `json:"complainant_position_lon"`
	ComplainedCity             string    `json:"complained_city"`
	ComplainedDistrict         string    `json:"complained_district"`
	ComplainedSubDistrict      string    `json:"complained_sub_district"`
	ComplainedAddress          string    `json:"complained_address"`
	ComplainedGMapLink         string    `json:"complained_gmap_link"`
	CategoryID                 uuid.UUID `json:"category_id"`
	CategoryName               string    `json:"category_name"`
	SubCategoryID              uuid.UUID `json:"sub_category_id"`
	SubCategoryName            string    `json:"sub_category_name"`
	Status                     int64     `json:"status"`
	Reason                     string    `json:"reason"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}

type AduanDetailResponse struct {
	ID         uuid.UUID `json:"id"`
	ComplainID string    `json:"complain_id"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`

	Status struct {
		Status int64  `json:"status"`
		Detail string `json:"detail"`
		Reason string `json:"reason"`
	} `json:"status"`

	Complainant struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Email string `json:"email"`

		SocialMedia struct {
			ID   uuid.UUID `json:"id"`
			Name string    `json:"name"`
			Link string    `json:"link"`
		} `json:"social_media"`

		Location struct {
			ComplainedLocation struct {
				City        string `json:"city"`
				District    string `json:"district"`
				SubDistrict string `json:"sub_district"`
				Link        string `json:"link"`
				Detail      string `json:"detail"`
			} `json:"complained_location"`

			ComplainantCoordinates struct {
				Lat string `json:"lat"`
				Lon string `json:"lon"`
			} `json:"complainant_coordinates"`
		} `json:"location"`

		Category struct {
			ID   uuid.UUID `json:"id"`
			Name string    `json:"name"`

			SubCategory struct {
				ID   uuid.UUID `json:"id"`
				Name string    `json:"name"`
			} `json:"sub_category"`
		} `json:"category"`

		Title       string `json:"title"`
		Description string `json:"description"`
		Evidence    string `json:"evidence"`
	} `json:"complainant"`
}

type AduanTableResponse struct {
	ID              uuid.UUID `json:"id"`
	ComplainID      string    `json:"complain_id"`
	ComplainantName string    `json:"complainant_name"`
	CategoryName    string    `json:"category_name"`
	CreatedAt       time.Time `json:"created_at"`
	Status          int64     `json:"status"`
}

type AduanSummaryResponse struct {
	Total         int64 `json:"total_complain"`
	TotalWaiting  int64 `json:"total_waiting"`
	TotalVerified int64 `json:"total_verified"`
	TotalFailed   int64 `json:"total_failed"`
}

type UpdateStatusPayload struct {
	ID     uuid.UUID `json:"id"`
	Status int64     `json:"status"`
	Reason string    `json:"reason"`
}

type AduanUsecase interface {
	Fetch(query *RequestQuery) (res []AduanTableResponse, totalData int64, err error)
	FindById(id uuid.UUID) (res *AduanDetail, err error)
	FetchSummary() (res *AduanSummaryResponse, err error)
	UpdateStatus(p *UpdateStatusPayload) (res interface{}, err error)
}

type AduanRepository interface {
	Fetch(query *RequestQuery) (res []AduanTableResponse, totalData int64, err error)
	FindById(id uuid.UUID) (res *AduanDetail, err error)
	FetchSummary() (res *AduanSummaryResponse, err error)
	UpdateStatus(p *UpdateStatusPayload) (res interface{}, err error)
	IsAduanExist(id uuid.UUID) (result bool)
}
