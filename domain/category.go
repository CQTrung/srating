package domain

import (
	"context"
)

type Category struct {
	HardModel
	Name string `json:"name" gorm:"column:name" validate:"required"`
}

type CategoryService interface {
	CreateCategory(c context.Context, department *Category) error
	GetAllCategory(c context.Context, input GetAllCategoryRequest) (int64, int64, []*Category, error)
	GetCategoryDetail(c context.Context, id uint) (*Category, error)
	UpdateCategory(c context.Context, department *Category) error
	DeleteCategory(c context.Context, id uint) error
}
type CategoryRepository interface {
	CreateCategory(c context.Context, department *Category) error
	GetAllCategory(c context.Context, input GetAllCategoryRequest) (int64, int64, []*Category, error)
	GetCategoryDetail(c context.Context, id uint) (*Category, error)
	UpdateCategory(c context.Context, department *Category) error
	DeleteCategory(c context.Context, id uint) error
}
type SearchCategoryRequest struct {
	UserID    uint  `json:"user_id"`
	Level     Level `json:"level"`
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
	PaginationRequest
}
type GetAllCategoryRequest struct {
	UserID    uint  `json:"user_id"`
	Level     Level `json:"level"`
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
	PaginationRequest
}
