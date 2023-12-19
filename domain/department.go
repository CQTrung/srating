package domain

import (
	"context"
)

type Department struct {
	HardModel
	Name  string  `json:"name" form:"name"`
	Users []*User `json:"users" gorm:"foreignKey:DepartmentID"`
}

type DepartmentService interface {
	CreateDepartment(c context.Context, department *Department) error
	GetAllDepartment(c context.Context) ([]*Department, error)
	GetDepartmentDetail(c context.Context, id uint) (*Department, error)
	UpdateDepartment(c context.Context, department *Department) error
	DeleteDepartment(c context.Context, id uint) error
}
type DepartmentRepository interface {
	CreateDepartment(c context.Context, department *Department) error
	GetAllDepartment(c context.Context) ([]*Department, error)
	GetDepartmentDetail(c context.Context, id uint) (*Department, error)
	UpdateDepartment(c context.Context, department *Department) error
	DeleteDepartment(c context.Context, id uint) error
}
