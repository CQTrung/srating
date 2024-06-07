package repositories

import (
	"context"
	"time"

	"srating/domain"

	"gorm.io/gorm"
)

type feedbackRepository struct {
	Repository
}

func NewFeedbackRepository(db *gorm.DB) domain.FeedbackRepository {
	return &feedbackRepository{
		Repository{
			database: db,
		},
	}
}

func (r *feedbackRepository) CreateFeedback(c context.Context, feedback *domain.Feedback) error {
	db := r.GetDB(c)
	return db.Save(feedback).Error
}

func (r *feedbackRepository) GetAllFeedback(c context.Context, idLocation uint, input domain.GetAllFeedbackRequest) (int64, int64, []*domain.Feedback, error) {
	feedbacks := []*domain.Feedback{}
	total := int64(0)
	start := time.Unix(input.StartDate, 0)
	end := time.Unix(input.EndDate, 0)
	query := r.GetDB(c).Model(&domain.Feedback{}).Preload("FeedbackCategories").Preload("FeedbackCategories.Category").Preload("User")
	if input.UserID != 0 {
		query = query.Where("user_id = ?", input.UserID)
	}
	if input.Level != 0 {
		query = query.Where("level = ?", input.Level)
	}
	if input.StartDate != 0 {
		query = query.Where("created_at >= ?", start)
	}
	if input.EndDate != 0 {
		query = query.Where("created_at <= ?", end)
	}

	// Apply location filter if present
	query = query.Joins("JOIN users ON users.id = feedbacks.user_id")

	if idLocation != 0 {
		if idLocation == input.LocationID {
			query = query.Where("users.location_id = ?", idLocation)
		} else if input.LocationID != 0 {
			query = query.Where("users.location_id = 0")
		} else {
			query = query.Where("users.location_id = ?", idLocation)
		}
	} else if input.LocationID != 0 {
		query = query.Where("users.location_id = ?", input.LocationID)
	}

	if err := query.Count(&total).Error; err != nil {
		return 0, 0, nil, err
	}
	query = query.Scopes(r.Paginate(input.Page, input.Limit)).Order("updated_at desc")
	if err := query.Find(&feedbacks).Error; err != nil {
		return 0, 0, nil, err
	}
	totalCount := int64(0)
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, 0, nil, err
	}

	return total, totalCount, feedbacks, nil
}

func (r *feedbackRepository) GetFeedbackDetail(c context.Context, id uint) (*domain.Feedback, error) {
	var feedback domain.Feedback
	query := r.GetDB(c)
	if err := query.Where("id = ?", id).First(&feedback).Error; err != nil {
		return nil, err
	}
	return &feedback, nil
}

func (r *feedbackRepository) UpdateFeedback(c context.Context, feedback *domain.Feedback) error {
	db := r.GetDB(c)
	return db.Save(feedback).Error
}

func (r *feedbackRepository) DeleteFeedback(c context.Context, id uint) error {
	db := r.GetDB(c)
	return db.Delete(&domain.Feedback{}, id).Error
}

func (r *feedbackRepository) GetFeedbackByLevel(c context.Context, level int, userID uint) ([]*domain.Feedback, error) {
	feedbackList := []*domain.Feedback{}
	query := r.GetDB(c)
	if err := query.Model(&domain.Feedback{}).Where("level = ? AND user_id = ?", level, userID).Find(&feedbackList).Error; err != nil {
		return nil, err
	}

	return feedbackList, nil
}
