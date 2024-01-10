package repositories

import (
	"context"
	"time"

	"srating/domain"

	"gorm.io/gorm"
)

type feedbackCategoryRepository struct {
	Repository
}

func NewFeedbackCategoryRepository(db *gorm.DB) domain.FeedbackCategoryRepository {
	return &feedbackCategoryRepository{
		Repository{
			database: db,
		},
	}
}

func (r *feedbackCategoryRepository) CreateFeedbackCategory(c context.Context, feedbackCategory *domain.FeedbackCategory) error {
	db := r.GetDB(c)
	return db.Save(feedbackCategory).Error
}

func (r *feedbackCategoryRepository) GetAllFeedbackCategory(c context.Context, input domain.GetAllFeedbackCategoryRequest) (int64, int64, []*domain.FeedbackCategory, error) {
	feedbackCategorys := []*domain.FeedbackCategory{}
	total := int64(0)
	start := time.Unix(input.StartDate, 0)
	end := time.Unix(input.EndDate, 0)
	query := r.GetDB(c).Model(&domain.FeedbackCategory{})
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
	if err := query.Count(&total).Error; err != nil {
		return 0, 0, nil, err
	}
	query = query.Scopes(r.Paginate(input.Page, input.Limit)).Order("updated_at desc")
	if err := query.Find(&feedbackCategorys).Error; err != nil {
		return 0, 0, nil, err
	}
	totalCount := int64(0)
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, 0, nil, err
	}

	return total, totalCount, feedbackCategorys, nil
}

func (r *feedbackCategoryRepository) GetFeedbackCategoryDetail(c context.Context, id uint) (*domain.FeedbackCategory, error) {
	var feedbackCategory domain.FeedbackCategory
	query := r.GetDB(c)
	if err := query.Where("id = ?", id).First(&feedbackCategory).Error; err != nil {
		return nil, err
	}
	return &feedbackCategory, nil
}

func (r *feedbackCategoryRepository) UpdateFeedbackCategory(c context.Context, feedbackCategory *domain.FeedbackCategory) error {
	db := r.GetDB(c)
	return db.Save(feedbackCategory).Error
}

func (r *feedbackCategoryRepository) DeleteFeedbackCategory(c context.Context, id uint) error {
	db := r.GetDB(c)
	return db.Delete(&domain.FeedbackCategory{}, id).Error
}

func (r *feedbackCategoryRepository) CountFeedbackByType(c context.Context) ([]*domain.GetFeedbackCategoryByTypeResponse, error) {
	result := []*domain.GetFeedbackCategoryByTypeResponse{}
	query := r.GetDB(c)
	if err := query.Model(&domain.FeedbackCategory{}).Select("level, count(*)").Group("level").Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *feedbackCategoryRepository) GetFeedbackLevelByUserID(c context.Context, userID uint) ([]*domain.GetFeedbackCategoryByTypeResponse, error) {
	var result []*domain.GetFeedbackCategoryByTypeResponse
	query := r.GetDB(c)
	if err := query.Model(&domain.Feedback{}).Select("level, count(*)").Where("user_id = ?", userID).Group("level").Scan(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (r *feedbackCategoryRepository) GetTotalFeedbackCategory(c context.Context) (int64, error) {
	var total int64
	query := r.GetDB(c)
	if err := query.Model(&domain.FeedbackCategory{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
