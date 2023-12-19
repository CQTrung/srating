package services

import (
	"context"
	"time"

	"srating/domain"
	"srating/utils"
)

type dashboardService struct {
	feedbackService domain.FeedbackService
	userService     domain.UserService
	contextTimeout  time.Duration
}

func NewDashboardService(feedbackService domain.FeedbackService, userService domain.UserService, timeout time.Duration) domain.DashboardService {
	return &dashboardService{
		feedbackService: feedbackService,
		userService:     userService,
		contextTimeout:  timeout,
	}
}

func (u *dashboardService) Dashboard(c context.Context) (map[string]int64, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	resultByType, err := u.feedbackService.CountFeedbackByType(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get feedback by type")
		return nil, err
	}
	totalFeedback, err := u.feedbackService.GetTotalFeedBack(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get total feedback")
		return nil, err
	}
	mapResult := make(map[string]int64, 8)
	userByType, err := u.userService.CountUserByRole(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get user by type")
		return nil, err
	}
	totalField, err := u.userService.CountTotalField(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get total field")
		return nil, err
	}
	for k, v := range resultByType {
		switch k {
		case 1:
			mapResult["very_good"] = v
		case 2:
			mapResult["good"] = v
		case 3:
			mapResult["normal"] = v
		case 4:
			mapResult["bad"] = v
		}
	}
	mapResult["total"] = totalFeedback
	for k, v := range userByType {
		switch k {
		case "manager":
			mapResult["manager"] = v
		case "employee":
			mapResult["employee"] = v
		}
	}
	mapResult["fields"] = totalField
	return mapResult, nil
}
