package helpers

import (
	"alqinsidev/jsa-mini-project/aduan/domain"
	"fmt"
	"time"
)

func MapRawAduanDetail(data *domain.AduanDetail) *domain.AduanDetailResponse {
	result := &domain.AduanDetailResponse{}

	result.ID = data.ID
	result.ComplainID = data.ComplainID
	result.CreatedAt = data.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = data.UpdatedAt.Format(time.RFC3339)

	result.Status.Status = data.Status
	result.Status.Detail = MapStatusDetail(data.Status)
	result.Status.Reason = data.Reason

	result.Complainant.Name = data.ComplainantName
	result.Complainant.Phone = data.ComplainantPhone

	result.Complainant.SocialMedia.ID = data.ComplainantSocialMediaID
	result.Complainant.SocialMedia.Name = data.ComplainantSocialMediaName
	result.Complainant.SocialMedia.Link = data.ComplainantSocialMediaLink

	result.Complainant.Location.ComplainedLocation.City = data.ComplainedCity
	result.Complainant.Location.ComplainedLocation.District = data.ComplainedDistrict
	result.Complainant.Location.ComplainedLocation.SubDistrict = data.ComplainedSubDistrict
	result.Complainant.Location.ComplainedLocation.Link = data.ComplainedGMapLink
	result.Complainant.Location.ComplainedLocation.Detail = data.ComplainedAddress

	result.Complainant.Location.ComplainantCoordinates.Lat = fmt.Sprintf("%f", data.ComplainantPositionLat)
	result.Complainant.Location.ComplainantCoordinates.Lon = fmt.Sprintf("%f", data.ComplainantPositionLon)

	result.Complainant.Category.ID = data.CategoryID
	result.Complainant.Category.Name = data.CategoryName

	result.Complainant.Category.SubCategory.ID = data.SubCategoryID
	result.Complainant.Category.SubCategory.Name = data.SubCategoryName

	result.Complainant.Title = data.Title
	result.Complainant.Description = data.Description
	result.Complainant.Evidence = data.Evidence

	return result
}

func MapStatusDetail(status int64) string {
	switch status {
	case 1:
		return "Menunggu Verifikasi"
	case 2:
		return "Terverifikasi"
	case 90:
		return "Gagal Terverifikasi"
	default:
		return "UNKNOWN STATUS"
	}
}
