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

func (r *feedbackRepository) GetAllFeedback(c context.Context, input domain.GetAllFeedbackRequest) (int64, int64, []*domain.Feedback, error) {
	feedbacks := []*domain.Feedback{}
	total := int64(0)
	query := r.GetDB(c).Model(&domain.Feedback{})
	if input.UserID != 0 {
		query = query.Where("user_id = ?", input.UserID)
	}
	if input.Level != 0 {
		query = query.Where("level = ?", input.Level)
	}
	if input.StartDate != 0 {
		query = query.Where("created_at >= ?", input.StartDate)
	}
	if input.EndDate != 0 {
		query = query.Where("created_at <= ?", input.EndDate)
	}
	if err := query.Count(&total).Error; err != nil {
		return 0, 0, nil, err
	}
	if err := query.Count(&total).Error; err != nil {
		return 0, 0, nil, err
	}
	query = query.Scopes(r.Paginate(input.Page, input.Limit))
	if err := query.Order("updated_at desc").Find(&feedbacks).Error; err != nil {
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

func (f *feedbackRepository) SearchFeedback(c context.Context, input domain.SearchFeedbackRequest) (int64, int64, []*domain.Feedback, error) {
	feedbacks := []*domain.Feedback{}
	total := int64(0)
	query := f.GetDB(c).Model(&domain.Feedback{})
	if input.UserID != 0 {
		query = query.Where("user_id = ?", input.UserID)
	}
	if input.Level != 0 {
		query = query.Where("level = ?", input.Level)
	}
	if input.StartDate != 0 {
		query = query.Where("created_at >= ?", input.StartDate)
	}
	if input.EndDate != 0 {
		query = query.Where("created_at <= ?", input.EndDate)
	}
	if err := query.Count(&total).Error; err != nil {
		return 0, 0, nil, err
	}
	query = query.Scopes(f.Paginate(input.Page, input.Limit))
	if err := query.Order("updated_at desc").Find(&feedbacks).Error; err != nil {
		return 0, 0, nil, err
	}
	totalCount := int64(0)
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, 0, nil, err
	}

	return total, totalCount, feedbacks, nil
}


func (r *feedbackRepository) GetFeedbackByLevel(c context.Context, level int, userID uint) ([]*domain.Feedback, error) {
	feedbackList := []*domain.Feedback{}
    query := r.GetDB(c)
    if err := query.Model(&domain.Feedback{}).Where("level = ? AND user_id = ?", level,userID).Find(&feedbackList).Error; err != nil {
        return nil, err
    }
    
    return feedbackList, nil
}