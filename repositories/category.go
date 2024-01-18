package repositories

import (
	"context"
	"time"

	"srating/domain"

	"gorm.io/gorm"
)

type categoryRepository struct {
	Repository
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{
		Repository{
			database: db,
		},
	}
}

func (r *categoryRepository) CreateCategory(c context.Context, category *domain.Category) error {
	db := r.GetDB(c)
	return db.Save(category).Error
}

func (r *categoryRepository) GetAllCategory(c context.Context, input domain.GetAllCategoryRequest) (int64, int64, []*domain.Category, error) {
	categorys := []*domain.Category{}
	total := int64(0)
	start := time.Unix(input.StartDate, 0)
	end := time.Unix(input.EndDate, 0)
	query := r.GetDB(c).Model(&domain.Category{})
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
	if err := query.Find(&categorys).Error; err != nil {
		return 0, 0, nil, err
	}
	totalCount := int64(0)
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, 0, nil, err
	}

	return total, totalCount, categorys, nil
}

func (r *categoryRepository) GetCategoryDetail(c context.Context, id uint) (*domain.Category, error) {
	var category domain.Category
	query := r.GetDB(c)
	if err := query.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) UpdateCategory(c context.Context, category *domain.Category) error {
	db := r.GetDB(c)
	return db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(c context.Context, id uint) error {
	db := r.GetDB(c)
	db.Where("id = ?", id).Delete(&domain.Category{})
	return db.Delete(&domain.Category{}, id).Error
}
