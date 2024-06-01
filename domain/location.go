package domain

import (
	"context"
)

type Location struct {
	Model
	Code string `json:"code" gorm:"code"`
	Name string `json:"name" gorm:"name"`
}

type LocationService interface {
	CreateLocation(c context.Context, location *Location) error
	GetAllLocation(c context.Context, input GetAllLocationRequest) (int64, int64, []*Location, error)
	GetLocationDetail(c context.Context, id uint) (*Location, error)
	UpdateLocation(c context.Context, location *Location) error
	DeleteLocation(c context.Context, id uint) error
}
type LocationRepository interface {
	CreateLocation(c context.Context, location *Location) error
	GetAllLocation(c context.Context, input GetAllLocationRequest) (int64, int64, []*Location, error)
	GetLocationDetail(c context.Context, id uint) (*Location, error)
	UpdateLocation(c context.Context, location *Location) error
	DeleteLocation(c context.Context, id uint) error
}
type GetAllLocationRequest struct {
	PaginationRequest
}
