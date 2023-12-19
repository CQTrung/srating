package controllers

import (
	"errors"
	"strconv"

	"srating/bootstrap"
	"srating/domain"
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
		ID:           user.ID,
		Username:     user.Username,
		Phone:        user.Phone,
		Email:        user.Email,
		ShortName:    user.ShortName,
		FullName:     user.FullName,
		Field:        user.Field,
		Avatar:       user.Avatar,
		DepartmentID: user.DepartmentID,
		Role:         user.Role,
		Status:       user.Status,
	}
	rest.AssertNil(err)
	uc.SendData(c, response)
}

// ChangeStatus
// @Router /users [put]
// @Tags user
// @Summary Change status
// @Security ApiKeyAuth
// @Success 200 {object} string

func (uc *UserController) ChangeStatus(c *gin.Context) {
	rawUserID, _ := c.Get("x-user-id")
	userIDStr, ok := rawUserID.(string)
	if !ok {
		rest.AssertNil(errors.New("invalid user id"))
	}
	userID, err := strconv.Atoi(userIDStr)
	rest.AssertNil(err)
	_, err = uc.UserService.GetUserByID(c, uint(userID))
	rest.AssertNil(err)
	body := domain.ChangeStatusRequest{}
	rest.AssertNil(c.ShouldBindJSON(&body))
	if !body.Status.IsValid() {
		rest.AssertNil(errors.New("invalid status"))
	}
	println(uc.Env.RequestTimeout)
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
	users, err := uc.UserService.GetAllEmployee(c)
	rest.AssertNil(err)
	uc.SendData(c, users)
}
