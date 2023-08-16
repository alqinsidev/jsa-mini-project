package domain

import (
	"time"

	"github.com/google/uuid"
)

type Complainant struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	SocialMediaID   uuid.UUID `json:"social_media_id"`
	SocialMediaLink string    `json:"social_media_link"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
