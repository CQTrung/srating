package controllers

import (
	"srating/bootstrap"
	"srating/domain"
	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	DashboardService domain.DashboardService
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
	result, err := t.DashboardService.Dashboard(c)
	rest.AssertNil(err)
	t.SendData(c, result)
}
