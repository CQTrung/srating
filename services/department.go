package services

import (
	"context"
	"time"

	"srating/domain"
	"srating/utils"
)

type departmentService struct {
	departmentRepository domain.DepartmentRepository
	contextTimeout       time.Duration
}

func NewDepartmentService(departmentRepository domain.DepartmentRepository, timeout time.Duration) domain.DepartmentService {
	return &departmentService{
		departmentRepository: departmentRepository,
		contextTimeout:       timeout,
	}
}

func (u *departmentService) CreateDepartment(c context.Context, input *domain.Department) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.departmentRepository.CreateDepartment(ctx, input); err != nil {
		utils.LogError(err, "Failed to create department")
		return err
	}
	return nil
}

func (u *departmentService) GetAllDepartment(c context.Context, input domain.GetAllDepartmentRequest) (int64, int64, []*domain.Department, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if input.Limit < 0 {
		input.Limit = 10
	}
	if input.Page < 0 {
		input.Page = 1
	}
	total, totalCount, departments, err := u.departmentRepository.GetAllDepartment(ctx, input)
	if err != nil {
		utils.LogError(err, "Failed to get all department")
		return 0, 0, nil, err
	}
	return total, totalCount, departments, nil
}

func (u *departmentService) GetDepartmentDetail(c context.Context, id uint) (*domain.Department, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	department, err := u.departmentRepository.GetDepartmentDetail(ctx, id)
	if err != nil {
		utils.LogError(err, "Failed to get department detail")
		return nil, err
	}
	return department, nil
}

func (u *departmentService) UpdateDepartment(c context.Context, input *domain.Department) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.departmentRepository.UpdateDepartment(ctx, input); err != nil {
		utils.LogError(err, "Failed to update department")
		return err
	}
	return nil
}

func (u *departmentService) DeleteDepartment(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := u.departmentRepository.DeleteDepartment(ctx, id); err != nil {
		utils.LogError(err, "Failed to delete department")
		return err
	}
	return nil
}
