package services

import (
	"context"
	"errors"
	"time"

	"srating/domain"
	"srating/utils"
)

type userService struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserService(userRepository domain.UserRepository, timeout time.Duration) domain.UserService {
	return &userService{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userService) GetUserByID(c context.Context, id uint) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user, err := uu.userRepository.GetUserByID(ctx, id)
	if err != nil {
		utils.LogError(err, "Failed to retrieve user")
		return nil, err
	}
	return user, nil
}

func (uu *userService) ChangeStatus(c context.Context, id uint, status domain.Status) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	if err := uu.userRepository.ChangeStatus(ctx, id, status); err != nil {
		utils.LogError(err, "Failed to change status")
		return err
	}
	return nil
}

func (uu *userService) GetAllEmployee(c context.Context) ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	users, err := uu.userRepository.GetAllEmployee(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get all employee")
		return nil, err
	}
	return users, nil
}

func (uu *userService) CountUserByRole(c context.Context) (map[string]int64, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	users, err := uu.userRepository.CountUserByRole(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get all employee")
		return nil, err
	}
	mapResult := make(map[string]int64, 2)

	for _, user := range users {
		mapResult[user.Role] = int64(user.Count)
	}
	if _, ok := mapResult[string(domain.EmployeeRole)]; !ok {
		mapResult[string(domain.EmployeeRole)] = 0
	}
	if _, ok := mapResult[string(domain.ManagerRole)]; !ok {
		mapResult[string(domain.ManagerRole)] = 0
	}

	return mapResult, nil
}

func (uu *userService) CountTotalField(c context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	count, err := uu.userRepository.CountTotalField(ctx)
	if err != nil {
		utils.LogError(err, "Failed to get total field")
		return 0, err
	}
	return count, nil
}

func (uu *userService) UpdateEmployee(c context.Context, input *domain.UpdateUserRequest) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if input.Role != domain.EmployeeRole {
		err := errors.New("failed to update employee")
		utils.LogError(err, "Failed to update employee")
		return err
	}
	input.Role = domain.EmployeeRole
	user := &domain.User{
		HardModel: domain.HardModel{
			ID: input.ID,
		},
		Email:        input.Email,
		ShortName:    input.ShortName,
		FullName:     input.FullName,
		Field:        input.Field,
		DepartmentID: input.DepartmentID,
		Phone:        input.Phone,
		Role:         input.Role,
		Status:       input.Status,
	}

	if err := uu.userRepository.UpdateUser(ctx, user); err != nil {
		utils.LogError(err, "Failed to update employee")
		return err
	}
	return nil
}

func (uu *userService) DeleteEmployee(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	if err := uu.userRepository.DeleteEmployee(ctx, id); err != nil {
		utils.LogError(err, "Failed to delete employee")
		return err
	}
	return nil
}

func (uu *userService) CreateUser(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	if err := utils.Validate(user); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if user.Role != domain.EmployeeRole {
		err := errors.New("failed to create employee")
		utils.LogError(err, "Failed to create employee")
		return err
	}
	user.Role = domain.EmployeeRole
	err := uu.userRepository.CreateUser(ctx, user)
	if err != nil {
		utils.LogError(err, "Failed to create employee")
		return err
	}
	return nil
}
