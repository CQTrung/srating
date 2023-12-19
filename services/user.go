package services

import (
	"context"
	"time"

	"srating/domain"
	"srating/utils"

	"github.com/hibiken/asynq"
)

type userService struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserService(userRepository domain.UserRepository, asynqClient *asynq.Client, timeout time.Duration) domain.UserService {
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
