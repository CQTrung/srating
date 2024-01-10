package services

import (
	"context"
	"time"

	"srating/domain"
	"srating/utils"
)

type feedbackService struct {
	feedbackRepository      domain.FeedbackRepository
	feedbackCategoryService domain.FeedbackCategoryService
	contextTimeout          time.Duration
}

func NewFeedbackService(feedbackRepository domain.FeedbackRepository,
	feedbackCategoryService domain.FeedbackCategoryService,
	timeout time.Duration,
) domain.FeedbackService {
	return &feedbackService{
		feedbackRepository:      feedbackRepository,
		feedbackCategoryService: feedbackCategoryService,
		contextTimeout:          timeout,
	}
}

func (u *feedbackService) CreateFeedback(c context.Context, input *domain.Feedback) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.feedbackRepository.CreateFeedback(ctx, input); err != nil {
		utils.LogError(err, "Failed to create feedback")
		return err
	}
	return nil
}

func (u *feedbackService) GetAllFeedback(c context.Context, input domain.GetAllFeedbackRequest) (int64, int64, []*domain.Feedback, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if input.Limit < 0 {
		input.Limit = 10
	}
	if input.Page < 0 {
		input.Page = 1
	}
	total, totalCount, feedbacks, err := u.feedbackRepository.GetAllFeedback(ctx, input)
	if err != nil {
		utils.LogError(err, "Failed to get all feedback")
		return 0, 0, nil, err
	}
	return total, totalCount, feedbacks, nil
}

func (u *feedbackService) GetFeedbackDetail(c context.Context, id uint) (*domain.Feedback, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	feedback, err := u.feedbackRepository.GetFeedbackDetail(ctx, id)
	if err != nil {
		utils.LogError(err, "Failed to get feedback detail")
		return nil, err
	}
	return feedback, nil
}

func (u *feedbackService) GetFeedbackByLevel(c context.Context, level int, userID uint) ([]*domain.Feedback, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	feedbacks, err := u.feedbackRepository.GetFeedbackByLevel(ctx, level, userID)
	if err != nil {
		utils.LogError(err, "Failed to get feedback by level")
		return nil, err
	}
	return feedbacks, nil
}

func (u *feedbackService) GetFeedbackLevelByUserID(c context.Context, userID uint) (map[string]int64, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	result, err := u.feedbackCategoryService.GetFeedbackLevelByUserID(ctx, userID)
	if err != nil {
		utils.LogError(err, "Failed to get feedback level by user id")
		return nil, err
	}
	return result, nil
}
