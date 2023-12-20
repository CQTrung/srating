package domain

import (
	"context"
)

type Role string

const (
	EmployeeRole = Role("employee")
	AdminRole    = Role("admin")
	ManagerRole  = Role("manager")
)

func (o Role) IsValid() bool {
	return o == EmployeeRole ||
		o == AdminRole ||
		o == ManagerRole
}

type Status int

const (
	ActiveStatus   = Status(1)
	InActiveStatus = Status(0)
)

func (s Status) IsValid() bool {
	return s == ActiveStatus ||
		s == InActiveStatus
}

type User struct {
	HardModel
	Username     string      `json:"username" gorm:"column:username" validate:"required"`
	Password     string      `json:"password" gorm:"column:password" validate:"required"`
	Phone        string      `json:"phone" gorm:"column:phone" validate:"required"`
	Email        string      `json:"email" gorm:"column:email" validate:"required"`
	ShortName    string      `json:"short_name" gorm:"column:short_name" validate:"required"`
	FullName     string      `json:"full_name" gorm:"column:full_name" validate:"required"`
	Field        string      `json:"field" gorm:"column:field" validate:"required"`
	Avatar       *Media      `json:"media" gorm:"many2many:user_media;"`
	DepartmentID uint        `json:"department_id" gorm:"department_id" validate:"required"`
	Department   *Department `json:"department" gorm:"foreignKey:DepartmentID"`
	Feedbacks    []*Feedback `json:"feedbacks" gorm:"foreignKey:UserID"`
	Role         Role        `json:"role" gorm:"column:role"`
	Status       Status      `json:"status" gorm:"column:status"`
}

type UserService interface {
	GetUserByID(c context.Context, id uint) (*User, error)
	ChangeStatus(c context.Context, id uint, status Status) error
	GetAllEmployee(c context.Context) ([]*User, error)
	CountUserByRole(c context.Context) (map[string]int64, error)
	CountTotalField(c context.Context) (int64, error)
	UpdateEmployee(c context.Context, input *UpdateUserRequest) error
	DeleteEmployee(c context.Context, id uint) error
	CreateUser(c context.Context, user *User) error
}
type UserRepository interface {
	UpdateUser(c context.Context, user *User) error
	GetUserByID(c context.Context, id uint) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	ChangeStatus(c context.Context, id uint, status Status) error
	GetAllEmployee(c context.Context) ([]*User, error)
	CountUserByRole(c context.Context) ([]*GetUserByRoleResponse, error)
	CountTotalField(c context.Context) (int64, error)
	UpdateEmployee(c context.Context, user *User) error
	DeleteEmployee(c context.Context, id uint) error
	CreateUser(c context.Context, user *User) error
}
type GetUserByIDResponse struct {
	ID           uint        `json:"id"`
	Username     string      `json:"username"`
	Phone        string      `json:"phone"`
	Email        string      `json:"email"`
	ShortName    string      `json:"short_name"`
	FullName     string      `json:"full_name"`
	Field        string      `json:"field"`
	Avatar       *Media      `json:"media"`
	DepartmentID uint        `json:"department_id"`
	Department   *Department `json:"department"`
	Role         Role        `json:"role"`
	Status       Status      `json:"status"`
}

type GetUserByRoleResponse struct {
	Role  string `json:"role"`
	Count int    `json:"count"`
}
type CreateUserRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	ShortName    string `json:"short_name"`
	FullName     string `json:"full_name"`
	Field        string `json:"field"`
	DepartmentID uint   `json:"department_id"`
	Role         Role   `json:"role"`
	Status       Status `json:"status"`
}
type UpdateUserRequest struct {
	ID           uint   `json:"id"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	ShortName    string `json:"short_name"`
	FullName     string `json:"full_name"`
	Field        string `json:"field"`
	DepartmentID uint   `json:"department_id"`
	Role         Role   `json:"role"`
	Status       Status `json:"status"`
}
