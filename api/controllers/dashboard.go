package controllers

import (
	"errors"
	"srating/bootstrap"
	"srating/domain"
	"srating/x/rest"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	DashboardService domain.DashboardService
	UserService      domain.UserService
	Env              *bootstrap.Env
	*rest.JSONRender
}

// Dashboard
// @Router /dashboard [get]
// @Tags dashboard
// @Query body domain.Dashboard
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *DashboardController) Dashboard(c *gin.Context) {
	rawUserID, _ := c.Get("x-user-id")
	userIDStr, ok := rawUserID.(string)
	if !ok {
		rest.AssertNil(errors.New("invalid user id"))
	}
	userID, err := strconv.Atoi(userIDStr)
	rest.AssertNil(err)
	user, err := t.UserService.GetUserByID(c, uint(userID))
	rest.AssertNil(err)

	if user.Role != domain.AdminRole && user.Role != domain.ManagerRole {
		rest.AssertNil(errors.New("permission denied"))
	}
	result, err := t.DashboardService.Dashboard(c, user.LocationID)
	rest.AssertNil(err)
	t.SendData(c, result)
}
