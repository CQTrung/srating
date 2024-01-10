package controllers

import (
	"strconv"

	"srating/bootstrap"
	"srating/domain"
	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	DepartmentService domain.DepartmentService
	Env               *bootstrap.Env
	*rest.JSONRender
}

// CreateDepartment
// @Router /departments [post]
// @Tags department
// @Query body domain.Department
// @Param payload body domain.Department true "payload"
// @Summary Create department
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *DepartmentController) CreateDepartment(c *gin.Context) {
	input := &domain.Department{}
	rest.AssertNil(c.ShouldBindJSON(&input))
	err := t.DepartmentService.CreateDepartment(c, input)
	rest.AssertNil(err)
	t.Success(c)
}

// GetAllDepartment
// @Router /departments [get]
// @Tags department
// @Summary Get all department
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param user_id query int false "user_id"
// @Param level query int false "level"
// @Param start_date query int false "start_date"
// @Param end_date query int false "end_date"
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *DepartmentController) GetAllDepartment(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	input := domain.GetAllDepartmentRequest{
		PaginationRequest: domain.PaginationRequest{
			Limit: limit,
			Page:  page,
		},
	}
	total, totalCount, result, err := t.DepartmentService.GetAllDepartment(c, input)
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

// GetDepartmentDetail
// @Router /departments/:id [get]
// @Tags department
// @Summary Get department by detail
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *DepartmentController) GetDepartmentDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	rest.AssertNil(err)
	result, err := t.DepartmentService.GetDepartmentDetail(c, uint(id))
	rest.AssertNil(err)
	t.SendData(c, result)
}

// UpdateDepartment
// @Router /departments [put]
// @Tags department
// @Summary Update department
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *DepartmentController) UpdateDepartment(c *gin.Context) {
	input := &domain.Department{}
	rest.AssertNil(c.ShouldBindJSON(&input))
	err := t.DepartmentService.UpdateDepartment(c, input)
	rest.AssertNil(err)
	t.Success(c)
}

// DeleteDepartment
// @Router /departments/:id [delete]
// @Tags department
// @Summary Delete department
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *DepartmentController) DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	rest.AssertNil(err)
	err = t.DepartmentService.DeleteDepartment(c, uint(id))
	rest.AssertNil(err)
	t.Success(c)
}
