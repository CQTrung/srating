package repositories

import (
	"context"

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

func (r *feedbackRepository) GetAllFeedback(c context.Context) ([]*domain.Feedback, error) {
	var feedbacks []*domain.Feedback
	query := r.GetDB(c)
	if err := query.Order("updated_at desc").Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	return feedbacks, nil
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

func (r *feedbackRepository) CountFeedbackByType(c context.Context) ([]*domain.GetFeedbackByTypeResponse, error) {
	result := []*domain.GetFeedbackByTypeResponse{}
	query := r.GetDB(c)
	if err := query.Model(&domain.Feedback{}).Select("level, count(*)").Group("level").Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *feedbackRepository) GetTotalFeedBack(c context.Context) (int64, error) {
	var total int64
	query := r.GetDB(c)
	if err := query.Model(&domain.Feedback{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *feedbackRepository) GetFeedbackByLevel(c context.Context, level int, userID uint) ([]*domain.Feedback, error) {
	feedbackList := []*domain.Feedback{}
    query := r.GetDB(c)
    if err := query.Model(&domain.Feedback{}).Where("level = ? AND user_id = ?", level,userID).Find(&feedbackList).Error; err != nil {
        return nil, err
    }
    
    return feedbackList, nil
}