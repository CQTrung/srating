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
	err := uu.userRepository.ChangeStatus(ctx, id, status)
	if err != nil {
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
