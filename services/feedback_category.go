package services

import (
	"context"
	"time"

	"srating/domain"
	"srating/utils"
)

type feedbackCategoryService struct {
	feedbackCategoryRepository domain.FeedbackCategoryRepository
	contextTimeout             time.Duration
}

func NewFeedbackCategoryService(feedbackCategoryRepository domain.FeedbackCategoryRepository,
	timeout time.Duration,
) domain.FeedbackCategoryService {
	return &feedbackCategoryService{
		feedbackCategoryRepository: feedbackCategoryRepository,
		contextTimeout:             timeout,
	}
}

func (u *feedbackCategoryService) CreateFeedbackCategory(c context.Context, input *domain.FeedbackCategory) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.feedbackCategoryRepository.CreateFeedbackCategory(ctx, input); err != nil {
		utils.LogError(err, "Failed to create feedbackCategory")
		return err
	}

	return nil
}

func (u *feedbackCategoryService) GetAllFeedbackCategory(c context.Context, input domain.GetAllFeedbackCategoryRequest) (int64, int64, []*domain.FeedbackCategory, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if input.Limit < 0 {
		input.Limit = 10
	}
	if input.Page < 0 {
		input.Page = 1
	}
	total, totalCount, feedbackCategorys, err := u.feedbackCategoryRepository.GetAllFeedbackCategory(ctx, input)
	if err != nil {
		utils.LogError(err, "Failed to get all feedbackCategory")
		return 0, 0, nil, err
	}
	return total, totalCount, feedbackCategorys, nil
}

func (u *feedbackCategoryService) GetFeedbackCategoryDetail(c context.Context, id uint) (*domain.FeedbackCategory, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	feedbackCategory, err := u.feedbackCategoryRepository.GetFeedbackCategoryDetail(ctx, id)
	if err != nil {
		utils.LogError(err, "Failed to get feedbackCategory detail")
		return nil, err
	}
	return feedbackCategory, nil
}

func (u *feedbackCategoryService) UpdateFeedbackCategory(c context.Context, input *domain.FeedbackCategory) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.feedbackCategoryRepository.UpdateFeedbackCategory(ctx, input); err != nil {
		utils.LogError(err, "Failed to update feedbackCategory")
		return err
	}

	return nil
}

func (u *feedbackCategoryService) DeleteFeedbackCategory(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := u.feedbackCategoryRepository.DeleteFeedbackCategory(ctx, id); err != nil {
		utils.LogError(err, "Failed to delete feedbackCategory")
		return err
	}

	return nil
}

func (u *feedbackCategoryService) CountFeedbackByType(c context.Context) (map[int]int64, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	result, err := u.feedbackCategoryRepository.CountFeedbackByType(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get feedback by type")
		return nil, err
	}
	mapResult := make(map[int]int64, 4)
	for _, v := range result {
		mapResult[v.Level] = int64(v.Count)
	}
	for i := 1; i <= 4; i++ {
		if _, ok := mapResult[i]; !ok {
			mapResult[i] = 0
		}
	}
	return mapResult, nil
}

func (u *feedbackCategoryService) GetFeedbackLevelByUserID(c context.Context, userID uint) (map[string]int64, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	result, err := u.feedbackCategoryRepository.GetFeedbackLevelByUserID(ctx, userID)
	if err != nil {
		utils.LogError(err, "Failed to get feedback level by user id")
		return nil, err
	}
	mapTempResult := make(map[int]int64, 4)
	for _, v := range result {
		mapTempResult[v.Level] = int64(v.Count)
	}
	for i := 1; i <= 4; i++ {
		if _, ok := mapTempResult[i]; !ok {
			mapTempResult[i] = 0
		}
	}
	mapResult := make(map[string]int64, 4)
	mapResult["very_good"] = mapTempResult[1]
	mapResult["good"] = mapTempResult[2]
	mapResult["normal"] = mapTempResult[3]
	mapResult["bad"] = mapTempResult[4]
	return mapResult, nil
}

func (u *feedbackCategoryService) GetTotalFeedbackCategory(c context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	total, err := u.feedbackCategoryRepository.GetTotalFeedbackCategory(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get total feedback")
		return 0, err
	}
	return total, nil
}
