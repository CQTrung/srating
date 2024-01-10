package services

import (
	"context"
	"time"

	"srating/domain"
	"srating/utils"
)

type categoryService struct {
	categoryRepository domain.CategoryRepository
	contextTimeout     time.Duration
}

func NewCategoryService(categoryRepository domain.CategoryRepository, timeout time.Duration) domain.CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
		contextTimeout:     timeout,
	}
}

func (u *categoryService) CreateCategory(c context.Context, input *domain.Category) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.categoryRepository.CreateCategory(ctx, input); err != nil {
		utils.LogError(err, "Failed to create category")
		return err
	}
	return nil
}

func (u *categoryService) GetAllCategory(c context.Context, input domain.GetAllCategoryRequest) (int64, int64, []*domain.Category, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if input.Limit < 0 {
		input.Limit = 10
	}
	if input.Page < 0 {
		input.Page = 1
	}
	total, totalCount, categorys, err := u.categoryRepository.GetAllCategory(ctx, input)
	if err != nil {
		utils.LogError(err, "Failed to get all category")
		return 0, 0, nil, err
	}
	return total, totalCount, categorys, nil
}

func (u *categoryService) GetCategoryDetail(c context.Context, id uint) (*domain.Category, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	category, err := u.categoryRepository.GetCategoryDetail(ctx, id)
	if err != nil {
		utils.LogError(err, "Failed to get category detail")
		return nil, err
	}
	return category, nil
}

func (u *categoryService) UpdateCategory(c context.Context, input *domain.Category) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.categoryRepository.UpdateCategory(ctx, input); err != nil {
		utils.LogError(err, "Failed to update category")
		return err
	}
	return nil
}

func (u *categoryService) DeleteCategory(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := u.categoryRepository.DeleteCategory(ctx, id); err != nil {
		utils.LogError(err, "Failed to delete category")
		return err
	}
	return nil
}
