package controllers

import (
	"strconv"

	"srating/bootstrap"
	"srating/domain"
	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService domain.CategoryService
	Env             *bootstrap.Env
	*rest.JSONRender
}

// CreateCategory
// @Router /categories [post]
// @Tags category
// @Query body domain.Category
// @Param payload body domain.Category true "payload"
// @Summary Create category
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *CategoryController) CreateCategory(c *gin.Context) {
	input := &domain.Category{}
	rest.AssertNil(c.ShouldBindJSON(&input))
	err := t.CategoryService.CreateCategory(c, input)
	rest.AssertNil(err)
	t.Success(c)
}

// GetAllCategory
// @Router /categories [get]
// @Tags category
// @Summary Get all category
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param user_id query int false "user_id"
// @Param level query int false "level"
// @Param start_date query int false "start_date"
// @Param end_date query int false "end_date"
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *CategoryController) GetAllCategory(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	userID, _ := strconv.Atoi(c.Query("user_id"))
	level, _ := strconv.Atoi(c.Query("level"))
	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))
	input := domain.GetAllCategoryRequest{
		UserID:    uint(userID),
		Level:     domain.Level(level),
		StartDate: int64(startDate),
		EndDate:   int64(endDate),
		PaginationRequest: domain.PaginationRequest{
			Limit: limit,
			Page:  page,
		},
	}
	total, totalCount, result, err := t.CategoryService.GetAllCategory(c, input)
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

// GetCategoryDetail
// @Router /categories/:id [get]
// @Tags category
// @Summary Get category by detail
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *CategoryController) GetCategoryDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	rest.AssertNil(err)
	result, err := t.CategoryService.GetCategoryDetail(c, uint(id))
	rest.AssertNil(err)
	t.SendData(c, result)
}

// UpdateCategory
// @Router /categories [put]
// @Tags category
// @Summary Update category
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *CategoryController) UpdateCategory(c *gin.Context) {
	input := &domain.Category{}
	rest.AssertNil(c.ShouldBindJSON(&input))
	err := t.CategoryService.UpdateCategory(c, input)
	rest.AssertNil(err)
	t.Success(c)
}

// DeleteCategory
// @Router /categories/:id [delete]
// @Tags category
// @Summary Delete category
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *CategoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	rest.AssertNil(err)
	err = t.CategoryService.DeleteCategory(c, uint(id))
	rest.AssertNil(err)
	t.Success(c)
}
