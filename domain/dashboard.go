package domain

import (
	"context"
)

type DashboardService interface {
	Dashboard(c context.Context) (map[string]int64, error)
}
