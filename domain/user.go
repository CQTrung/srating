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
	Username     string      `json:"username" gorm:"column:username;unique" validate:"required"`
	Password     string      `json:"password" gorm:"column:password" validate:"required"`
	Phone        string      `json:"phone" gorm:"column:phone" validate:"required"`
	Email        string      `json:"email" gorm:"column:email;unique" validate:"required"`
	ShortName    string      `json:"short_name" gorm:"column:short_name"`
	FullName     string      `json:"full_name" gorm:"column:full_name" validate:"required"`
	Field        string      `json:"field" gorm:"column:field" validate:"required"`
	MediaID      uint        `json:"media_id" gorm:"column:media_id;default:1"`
	Avatar       *Media      `json:"media" gorm:"foreignKey:MediaID"`
	DepartmentID uint        `json:"department_id" gorm:"department_id"`
	Department   *Department `json:"department" gorm:"foreignKey:DepartmentID"`
	Counter      string      `json:"counter" gorm:"column:counter"`
	Feedbacks    []*Feedback `json:"feedbacks,omitempty" gorm:"foreignKey:UserID"`
	Role         Role        `json:"role" gorm:"column:role"`
	Status       Status      `json:"status" gorm:"column:status"`
	LocationID   uint        `json:"location_id" gorm:"column:location_id"`
	Location     *Location   `json:"location" gorm:"foreignKey:LocationID"`
}

type UserService interface {
	GetUserByID(c context.Context, id uint) (*User, error)
	ChangeStatus(c context.Context, id uint, status Status) error
	GetAllEmployee(c context.Context, idLocation uint, input GetAllUserRequest) (int64, int64, []*User, error)
	CountUserByRole(c context.Context, idLocation uint) (map[string]int64, error)
	CountTotalField(c context.Context, idLocation uint) (int64, error)
	UpdateEmployee(c context.Context, user *User) error
	DeleteEmployee(c context.Context, id uint) error
	CreateUser(c context.Context, user *User) error
	ChangePassword(c context.Context, id uint, oldPassword, newPassword string) error
	ResetPassword(c context.Context, id uint) error
	AssignToLocation(c context.Context, input AssignLocationRequest) error
}
type UserRepository interface {
	UpdateUser(c context.Context, user *User) error
	GetUserByID(c context.Context, id uint) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	ChangeStatus(c context.Context, id uint, status Status) error
	GetAllEmployee(c context.Context, idLocation uint, input GetAllUserRequest) (int64, int64, []*User, error)
	CountUserByRole(c context.Context, idLocation uint) ([]*GetUserByRoleResponse, error)
	CountTotalField(c context.Context, idLocation uint) (int64, error)
	DeleteEmployee(c context.Context, id uint) error
	CreateUser(c context.Context, user *User) error
	ChangePassword(c context.Context, id uint, oldPassword, newPassword string) error
	ResetPassword(c context.Context, id uint, newPassword string) error
	AssignToLocation(c context.Context, input AssignLocationRequest) error
}
type GetUserByIDResponse struct {
	ID           uint        `json:"id"`
	Username     string      `json:"username"`
	Phone        string      `json:"phone"`
	Email        string      `json:"email"`
	ShortName    string      `json:"short_name"`
	FullName     string      `json:"full_name"`
	Field        string      `json:"field"`
	Counter      string      `json:"counter"`
	Avatar       *Media      `json:"media"`
	DepartmentID uint        `json:"department_id"`
	Department   *Department `json:"department"`
	LocationID   uint        `json:"location_id"`
	Location     *Location   `json:"location"`
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
	Counter      string `json:"counter"`
	Field        string `json:"field"`
	Role         Role   `json:"role"`
	DepartmentID uint   `json:"department_id"`
	LocationID   uint   `json:"location_id"`
	Status       Status `json:"status"`
}
type UpdateUserRequest struct {
	ID           uint   `json:"id"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Counter      string `json:"counter"`
	ShortName    string `json:"short_name"`
	FullName     string `json:"full_name"`
	Field        string `json:"field"`
	DepartmentID uint   `json:"department_id"`
	LocationID   uint   `json:"location_id"`
	Role         Role   `json:"role"`
	Status       Status `json:"status"`
}
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
type ResetPasswordRequest struct {
	ID uint `json:"id"`
}
type GetAllUserRequest struct {
	PaginationRequest
}

type AssignLocationRequest struct {
	Id         uint `json:"id"`
	LocationId uint `json:"locationId"`
}
