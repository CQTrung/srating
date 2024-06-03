package domain

import (
	"context"
)

type DashboardService interface {
	Dashboard(c context.Context,location uint) (map[string]int64, error)
}
