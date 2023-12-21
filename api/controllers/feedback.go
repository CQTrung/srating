package controllers

import (
	"strconv"

	"srating/bootstrap"
	"srating/domain"
	"srating/utils"
	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

type FeedbackController struct {
	FeedbackService domain.FeedbackService
	Env             *bootstrap.Env
	*rest.JSONRender
}

// CreateFeedback
// @Router /feedbacks [post]
// @Tags feedback
// @Query body domain.Feedback
// @Param payload body domain.Feedback true "payload"
// @Summary Create feedback
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *FeedbackController) CreateFeedback(c *gin.Context) {
	input := &domain.Feedback{}
	rest.AssertNil(c.ShouldBindJSON(&input))
	err := t.FeedbackService.CreateFeedback(c, input)
	rest.AssertNil(err)
	t.Success(c)
}

// GetAllFeedback
// @Router /feedbacks [get]
// @Tags feedback
// @Summary Get all feedback
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param user_id query int false "user_id"
// @Param level query int false "level"
// @Param start_date query int false "start_date"
// @Param end_date query int false "end_date"
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *FeedbackController) GetAllFeedback(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	userID, _ := strconv.Atoi(c.Query("user_id"))
	level, _ := strconv.Atoi(c.Query("level"))
	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))
	input := domain.GetAllFeedbackRequest{
		UserID:    uint(userID),
		Level:     domain.Level(level),
		StartDate: int64(startDate),
		EndDate:   int64(endDate),
		PaginationRequest: domain.PaginationRequest{
			Limit: limit,
			Page:  page,
		},
	}
	total, totalCount, result, err := t.FeedbackService.GetAllFeedback(c, input)
	rest.AssertNil(err)
	t.SendCustomData(c, map[string]interface{}{
		"status":     "success",
		"data":       result,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalCount": totalCount,
	},
	)
}

// GetFeedbackDetail
// @Router /feedbacks/:id [get]
// @Tags feedback
// @Summary Get feedback by detail
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *FeedbackController) GetFeedbackDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	rest.AssertNil(err)
	result, err := t.FeedbackService.GetFeedbackDetail(c, uint(id))
	rest.AssertNil(err)
	t.SendData(c, result)
}

// SearchFeedback
// @Router /feedbacks/search [post]
// @Tags feedback
// @Summary Search feedback
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *FeedbackController) SearchFeedback(c *gin.Context) {
	input := domain.SearchFeedbackRequest{}
	rest.AssertNil(c.ShouldBindJSON(&input))
	total, totalCount, result, err := t.FeedbackService.SearchFeedback(c, input)
	rest.AssertNil(err)
	t.SendCustomData(c, map[string]interface{}{
		"status":     "success",
		"data":       result,
		"total":      total,
		"totalCount": totalCount,
	},
	)
}

// GetFeedbackByLevel
// @Router /feedbacks/level
// @Tags feedback
// @Summary Get feedback by level
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *FeedbackController) GetFeedbackByLevel(c *gin.Context) {
	level, err := strconv.Atoi(c.Query("level"))
	rest.AssertNil(err)
	userID, err := utils.GetUserIDFromContext(c)
	rest.AssertNil(err)
	result, err := t.FeedbackService.GetFeedbackByLevel(c, level, userID)
	rest.AssertNil(err)
	t.SendData(c, result)
}
