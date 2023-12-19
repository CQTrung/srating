package domain

import (
	"context"
)

type Media struct {
	Model
	URL string `gorm:"column:url" json:"url"`
	// Type     string `gorm:"column:type" json:"type"`
	FileName string `gorm:"column:filename" json:"filename"`
}
type MediaRepository interface {
	Upload(ctx context.Context, media *Media) error
	GetAll(ctx context.Context) ([]*Media, error)
}
type MediaService interface {
	Upload(ctx context.Context, input []*UploadFileInput) ([]*Media, error)
	GetAll(ctx context.Context) ([]*Media, error)
}
