package services

import (
	"context"
	"time"

	"srating/domain"
	"srating/utils"
)

type feedbackService struct {
	feedbackRepository domain.FeedbackRepository
	contextTimeout     time.Duration
}

func NewFeedbackService(feedbackRepository domain.FeedbackRepository, timeout time.Duration) domain.FeedbackService {
	return &feedbackService{
		feedbackRepository: feedbackRepository,
		contextTimeout:     timeout,
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

func (u *feedbackService) GetAllFeedback(c context.Context) ([]*domain.Feedback, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	feedbacks, err := u.feedbackRepository.GetAllFeedback(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get all feedback")
		return nil, err
	}
	return feedbacks, nil
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

func (u *feedbackService) UpdateFeedback(c context.Context, input *domain.Feedback) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.feedbackRepository.UpdateFeedback(ctx, input); err != nil {
		utils.LogError(err, "Failed to update feedback")
		return err
	}
	return nil
}

func (u *feedbackService) DeleteFeedback(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := u.feedbackRepository.DeleteFeedback(ctx, id); err != nil {
		utils.LogError(err, "Failed to delete feedback")
		return err
	}
	return nil
}

func (u *feedbackService) CountFeedbackByType(c context.Context) (map[int]int64, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	result, err := u.feedbackRepository.CountFeedbackByType(ctx)
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

func (u *feedbackService) GetTotalFeedBack(c context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	total, err := u.feedbackRepository.GetTotalFeedBack(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get total feedback")
		return 0, err
	}
	return total, nil
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
