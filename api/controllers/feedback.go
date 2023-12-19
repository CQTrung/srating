package controllers

import (
	"strconv"

	"srating/bootstrap"
	"srating/domain"
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
	rest.AssertNil(c.ShouldBindJSON(input))
	err := t.FeedbackService.CreateFeedback(c, input)
	rest.AssertNil(err)
	t.Success(c)
}

// GetAllFeedback
// @Router /feedbacks [get]
// @Tags feedback
// @Summary Get all feedback
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *FeedbackController) GetAllFeedback(c *gin.Context) {
	result, err := t.FeedbackService.GetAllFeedback(c)
	rest.AssertNil(err)
	t.SendData(c, result)
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

// UpdateFeedback
// @Router /feedbacks [put]
// @Tags feedback
// @Param payload body domain.Feedback true "payload"
// @Summary Update feedback
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *FeedbackController) UpdateFeedback(c *gin.Context) {
	input := &domain.Feedback{}
	rest.AssertNil(c.ShouldBindJSON(input))
	err := t.FeedbackService.UpdateFeedback(c, input)
	rest.AssertNil(err)
	t.Success(c)
}

// // DeleteFeedback
// // @Router /feedbacks/:id [delete]
// // @Tags feedback
// // @Summary Delete feedback
// // @Security ApiKeyAuth
// // @Success 200 {object} string
// func (t *FeedbackController) DeleteFeedback(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	rest.AssertNil(err)
// 	err = t.FeedbackService.DeleteFeedback(c, uint(id))
// 	rest.AssertNil(err)
// 	t.Success(c)
// }
