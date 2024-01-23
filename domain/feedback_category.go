package domain

import (
	"context"
)

type Level int

const (
	VeryGood = Level(1)
	Good     = Level(2)
	Normal   = Level(3)
	Bad      = Level(4)
)

type FeedbackCategory struct {
	HardModel
	CategoryID uint      `json:"category_id" gorm:"column:category_id" validate:"required"`
	FeedbackID uint      `json:"feedback_id" gorm:"column:feedback_id"`
	Category   *Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Level      string    `json:"level" gorm:"column:level" validate:"required"`
	Note       string    `json:"note" gorm:"column:note"`
}
type FeedbackCategoryService interface {
	CreateFeedbackCategory(c context.Context, feedbackCategory *FeedbackCategory) error
	GetAllFeedbackCategory(c context.Context, input GetAllFeedbackCategoryRequest) (int64, int64, []*FeedbackCategory, error)
	GetFeedbackCategoryDetail(c context.Context, id uint) (*FeedbackCategory, error)
	UpdateFeedbackCategory(c context.Context, feedbackCategory *FeedbackCategory) error
	DeleteFeedbackCategory(c context.Context, id uint) error
	CountFeedbackByType(c context.Context) (map[int]int64, error)
	GetFeedbackLevelByUserID(c context.Context, userID uint) (map[string]int64, error)
	GetTotalFeedbackCategory(c context.Context) (int64, error)
}
type FeedbackCategoryRepository interface {
	CreateFeedbackCategory(c context.Context, feedbackCategory *FeedbackCategory) error
	GetAllFeedbackCategory(c context.Context, input GetAllFeedbackCategoryRequest) (int64, int64, []*FeedbackCategory, error)
	GetFeedbackCategoryDetail(c context.Context, id uint) (*FeedbackCategory, error)
	UpdateFeedbackCategory(c context.Context, feedbackCategory *FeedbackCategory) error
	DeleteFeedbackCategory(c context.Context, id uint) error
	CountFeedbackByType(c context.Context) ([]*GetFeedbackCategoryByTypeResponse, error)
	GetFeedbackLevelByUserID(c context.Context, userID uint) ([]*GetFeedbackCategoryByTypeResponse, error)
	GetTotalFeedbackCategory(c context.Context) (int64, error)
}
type GetFeedbackCategoryByTypeResponse struct {
	Level int `json:"level"`
	Count int `json:"count"`
}
type GetAllFeedbackCategoryResponse struct {
	FeedbackCategorys []*FeedbackCategory `json:"feedbacks"`
	Total             int64               `json:"total"`
	Page              int                 `json:"page"`
	Limit             int                 `json:"limit"`
}
type SearchFeedbackCategoryRequest struct {
	UserID    uint  `json:"user_id"`
	Level     Level `json:"level"`
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
	PaginationRequest
}
type GetAllFeedbackCategoryRequest struct {
	UserID    uint  `json:"user_id"`
	Level     Level `json:"level"`
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
	PaginationRequest
}
