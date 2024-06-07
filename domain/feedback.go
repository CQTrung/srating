package domain

import (
	"context"
	"time"
)

type Feedback struct {
	HardModel
	UserID             uint                `json:"user_id" gorm:"column:user_id" validate:"required"`
	User               *User               `json:"user" gorm:"foreignKey:UserID"`
	Level              Level               `json:"level"`
	Note               string              `json:"note" gorm:"column:note"`
	FeedbackCategories []*FeedbackCategory `json:"feedback_categories" gorm:"foreignkey:FeedbackID"`
}

type FeedbackReponse struct {
	ID                 uint                `json:"id"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	UserID             uint                `json:"user_id"`
	FullName           string              `json:"full_name"`
	LocationId         string              `json:"location_id"`
	Level              Level               `json:"level"`
	Note               string              `json:"note"`
	FeedbackCategories []*FeedbackCategory `json:"feedback_categories" gorm:"foreignkey:FeedbackID"`
}

type FeedbackService interface {
	CreateFeedback(c context.Context, department *Feedback) error
	CreateFeedbackV2(c context.Context, department *Feedback) error
	GetAllFeedback(c context.Context, idLocation uint, input GetAllFeedbackRequest) (int64, int64, []*FeedbackReponse, error)
	GetFeedbackDetail(c context.Context, id uint) (*Feedback, error)
	GetFeedbackByLevel(c context.Context, level int, userID uint) ([]*Feedback, error)
	GetFeedbackLevelByUserID(c context.Context, userID uint) (map[string]int64, error)
}
type FeedbackRepository interface {
	CreateFeedback(c context.Context, department *Feedback) error
	GetAllFeedback(c context.Context, idLocation uint, input GetAllFeedbackRequest) (int64, int64, []*FeedbackReponse, error)
	GetFeedbackDetail(c context.Context, id uint) (*Feedback, error)
	GetFeedbackByLevel(c context.Context, level int, userID uint) ([]*Feedback, error)
	Transaction(ctx context.Context, callback func(ctx context.Context) error) error
}

type GetAllFeedbackResponse struct {
	Feedbacks []*Feedback `json:"feedbacks"`
	Total     int64       `json:"total"`
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
}
type SearchFeedbackRequest struct {
	UserID    uint  `json:"user_id"`
	Level     Level `json:"level"`
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
	PaginationRequest
}
type GetAllFeedbackRequest struct {
	UserID     uint  `json:"user_id"`
	LocationID uint  `json:"location_id"`
	Level      Level `json:"level"`
	StartDate  int64 `json:"start_date"`
	EndDate    int64 `json:"end_date"`
	PaginationRequest
}
