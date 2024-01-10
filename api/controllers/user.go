package controllers

import (
	"errors"
	"strconv"

	"srating/bootstrap"
	"srating/domain"
	"srating/utils"
	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService domain.UserService
	Env         *bootstrap.Env
	*rest.JSONRender
}

// GetUserDetail
// @Router /users [get]
// @Tags user
// @Summary Get user detail
// @Security ApiKeyAuth
// @Success 200 {object} string
func (uc *UserController) GetUserDetail(c *gin.Context) {
	rawUserID, _ := c.Get("x-user-id")
	userIDStr, ok := rawUserID.(string)
	if !ok {
		rest.AssertNil(errors.New("invalid user id"))
	}
	userID, err := strconv.Atoi(userIDStr)
	rest.AssertNil(err)
	user, err := uc.UserService.GetUserByID(c, uint(userID))
	response := &domain.GetUserByIDResponse{
		ID:         user.ID,
		Username:   user.Username,
		Phone:      user.Phone,
		Email:      user.Email,
		ShortName:  user.ShortName,
		FullName:   user.FullName,
		Field:      user.Field,
		Counter:    user.Counter,
		Avatar:     user.Avatar,
		Department: user.Department,
		Role:       user.Role,
		Status:     user.Status,
	}
	rest.AssertNil(err)
	uc.SendData(c, response)
}

// ChangeStatus
// @Router /users/status [put]
// @Tags user
// @Summary Change status
// @Security ApiKeyAuth
// @Success 200 {object} string
func (uc *UserController) ChangeStatus(c *gin.Context) {
	rawUserID, _ := c.Get("x-user-id")
	userIDStr, ok := rawUserID.(string)
	if !ok {
		rest.AssertNil(errors.New("invalid User ID"))
	}
	userID, err := strconv.Atoi(userIDStr)
	rest.AssertNil(err)
	_, err = uc.UserService.GetUserByID(c, uint(userID))
	rest.AssertNil(err)
	var body *struct {
		Status domain.Status `json:"status"`
	}
	utils.LogError(c.ShouldBindJSON(&body), "?")
	if !body.Status.IsValid() {
		rest.AssertNil(errors.New("invalid status"))
	}
	err = uc.UserService.ChangeStatus(c, uint(userID), body.Status)
	rest.AssertNil(err)
	uc.Success(c)
}

// GetAllEmployee
// @Router /users/employees [get]
// @Tags user
// @Summary Get all employee
// @Security ApiKeyAuth
// @Success 200 {object} string
func (uc *UserController) GetAllEmployee(c *gin.Context) {
	rawUserID, _ := c.Get("x-user-id")
	userIDStr, ok := rawUserID.(string)
	if !ok {
		rest.AssertNil(errors.New("invalid user id"))
	}
	userID, err := strconv.Atoi(userIDStr)
	rest.AssertNil(err)
	user, err := uc.UserService.GetUserByID(c, uint(userID))
	rest.AssertNil(err)
	if user.Role != domain.AdminRole {
		rest.AssertNil(errors.New("permission denied"))
	}
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	params := domain.GetAllUserRequest{
		PaginationRequest: domain.PaginationRequest{
			Limit: limit,
			Page:  page,
		},
	}
	total, totalCount, users, err := uc.UserService.GetAllEmployee(c, params)
	rest.AssertNil(err)
	uc.SendCustomData(c, map[string]interface{}{
		"status":     "success",
		"data":       users,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalCount": totalCount,
	})
}

// UpdateEmployee
// @Router /users/employees [put]
// @Tags user
// @Param payload body domain.UpdateUserRequest true "payload"
// @Summary Update employee
// @Security ApiKeyAuth
// @Success 200 {object} string
func (uc *UserController) UpdateEmployee(c *gin.Context) {
	body := &domain.User{}
	rest.AssertNil(c.ShouldBindJSON(&body))
	err := uc.UserService.UpdateEmployee(c, body)
	rest.AssertNil(err)
	uc.Success(c)
}

// DeleteEmployee
// @Router /users/employees/:id [delete]
// @Tags user
// @Summary Delete employee
// @Security ApiKeyAuth
// @Success 200 {object} string
func (uc *UserController) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	rest.AssertNil(err)
	err = uc.UserService.DeleteEmployee(c, uint(id))
	rest.AssertNil(err)
	uc.Success(c)
}

// CreateUserByAdmin
// @Router /users [post]
// @Tags user
// @Summary Create user by admin
// @Param payload body domain.CreateUserRequest true "payload"
// @Security ApiKeyAuth
// @Success 200 {object} string
func (uc *UserController) CreateUser(c *gin.Context) {
	body := &domain.User{}
	rest.AssertNil(c.ShouldBindJSON(&body))
	err := uc.UserService.CreateUser(c, body)
	rest.AssertNil(err)
	uc.Success(c)
}

// ChangePassword
// @Router /users/change-password [post]
// @Tags user
// @Param payload body domain.ChangePasswordRequest true "payload"
// @Summary Change password
// @Security ApiKeyAuth
// @Success 200 {object} string
func (uc *UserController) ChangePassword(c *gin.Context) {
	rawUserID, _ := c.Get("x-user-id")
	userIDStr, ok := rawUserID.(string)
	if !ok {
		rest.AssertNil(errors.New("invalid user id"))
	}
	userID, err := strconv.Atoi(userIDStr)
	rest.AssertNil(err)
	body := &domain.ChangePasswordRequest{}
	err = c.ShouldBindJSON(&body)
	rest.AssertNil(err)
	err = uc.UserService.ChangePassword(c, uint(userID), body.OldPassword, body.NewPassword)
	// rest.AssertNil(err)
	rest.AssertNil(err)
	uc.Success(c)
}

// ResetPassword
// @Router /users/reset-password [post]
// @Tags user
// @Param payload body domain.ResetPasswordRequest true "payload"
// @Summary Reset password
// @Security ApiKeyAuth
// @Success 200 {object} string
func (uc *UserController) ResetPassword(c *gin.Context) {
	body := &domain.ResetPasswordRequest{}
	err := c.ShouldBindJSON(&body)
	rest.AssertNil(err)
	err = uc.UserService.ResetPassword(c, body.ID)
	rest.AssertNil(err)
	uc.Success(c)
}
