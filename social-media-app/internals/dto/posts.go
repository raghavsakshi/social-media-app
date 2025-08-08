package dto

import (
	"github.com/google/uuid"
	"time"
)

type PostCreate struct {

	Content string `json:"content" validate:"required,max=2000"`
	// UserID  uuid.UUID `json:"user_id" validate:"required"`
}
type Posts struct {
	ID        uuid.UUID  ` json:"id"`
	Content string     `json:"content"`
	UserID uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Post struct {
			ID        uuid.UUID  `json:"id"`
	Content string     `json:"content"`
	UserID uuid.UUID `json:"user_id"`
	CreatedAt *time.Time   `json:"created_at,omitempty"`
	UpdatedAt *time.Time   `json:"updated_at,omitempty"`
}