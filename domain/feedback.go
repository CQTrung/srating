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

type Feedback struct {
	HardModel
	UserID uint   `json:"user_id" gorm:"column:user_id" validate:"required"`
	Level  Level  `json:"level" gorm:"column:level" validate:"required"`
	Note   string `json:"note" gorm:"column:note"`
}

type FeedbackService interface {
	CreateFeedback(c context.Context, department *Feedback) error
	GetAllFeedback(c context.Context) ([]*Feedback, error)
	GetFeedbackDetail(c context.Context, id uint) (*Feedback, error)
	UpdateFeedback(c context.Context, department *Feedback) error
	DeleteFeedback(c context.Context, id uint) error
	CountFeedbackByType(c context.Context) (map[int]int64, error)
	GetTotalFeedBack(c context.Context) (int64, error)
	GetFeedbackByLevel(c context.Context, level int, userID uint) ([]*Feedback, error)
}
type FeedbackRepository interface {
	CreateFeedback(c context.Context, department *Feedback) error
	GetAllFeedback(c context.Context) ([]*Feedback, error)
	GetFeedbackDetail(c context.Context, id uint) (*Feedback, error)
	UpdateFeedback(c context.Context, department *Feedback) error
	DeleteFeedback(c context.Context, id uint) error
	CountFeedbackByType(c context.Context) ([]*GetFeedbackByTypeResponse, error)
	GetTotalFeedBack(c context.Context) (int64, error)
	GetFeedbackByLevel(c context.Context, level int, userID uint) ([]*Feedback, error)
}
type GetFeedbackByTypeResponse struct {
	Level int `json:"level"`
	Count int `json:"count"`
}
