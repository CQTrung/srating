package services

import (
	"context"
	"time"

	"srating/domain"
	"srating/utils"
	"srating/x/rest"
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

func (uu *userService) GetAllEmployee(c context.Context, idLocation uint, input domain.GetAllUserRequest) (int64, int64, []*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	total, totalCount, users, err := uu.userRepository.GetAllEmployee(ctx, idLocation, input)
	if err != nil {
		utils.LogError(err, "Failed to get all employee")
		return 0, 0, nil, err
	}
	return total, totalCount, users, nil
}

func (uu *userService) CountUserByRole(c context.Context,locationId uint) (map[string]int64, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	users, err := uu.userRepository.CountUserByRole(ctx,locationId)
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

func (uu *userService) CountTotalField(c context.Context,locationId uint) (int64, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	count, err := uu.userRepository.CountTotalField(ctx,locationId)
	if err != nil {
		utils.LogError(err, "Failed to get total field")
		return 0, err
	}
	return count, nil
}

func (uu *userService) UpdateEmployee(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user.Role = domain.EmployeeRole
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
	phone, ok := utils.ValidatePhone(user.Phone)
	if !ok {
		return rest.ErrValidation
	}
	user.Phone = phone
	user.Role = domain.EmployeeRole
	password := "123456"
	password, err := utils.GenerateHashPassword(password)
	user.Password = password
	if err != nil {
		utils.LogError(err, "Invalid password")
		return err
	}
	if err := uu.userRepository.CreateUser(ctx, user); err != nil {
		utils.LogError(err, "Failed to create employee")
		return err
	}
	return nil
}

func (uu *userService) ChangePassword(c context.Context, id uint, oldPassword, newPassword string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user, err := uu.userRepository.GetUserByID(ctx, id)
	if err != nil {
		utils.LogError(err, "Failed to retrieve user")
		return err
	}

	if utils.CompareHashAndPassword(user.Password, oldPassword) != nil {
		return rest.ErrInvalidCredentials
	}
	oldPassword = user.Password
	newPassword, err = utils.GenerateHashPassword(newPassword)
	if err != nil {
		utils.LogError(err, "Invalid password")
		return err
	}
	if err := uu.userRepository.ChangePassword(ctx, id, oldPassword, newPassword); err != nil {
		utils.LogError(err, "Failed to change password")
		return err
	}
	return nil
}

func (uu *userService) ResetPassword(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	newPassword := "123456"
	newPassword, err := utils.GenerateHashPassword(newPassword)
	if err != nil {
		utils.LogError(err, "Invalid password")
		return err
	}
	if err := uu.userRepository.ResetPassword(ctx, id, newPassword); err != nil {
		utils.LogError(err, "Failed to change password")
		return err
	}
	return nil
}

func (uu *userService) AssignToLocation(c context.Context, input domain.AssignLocationRequest) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	// Directly call the repository method to assign the user to the location
	if err := uu.userRepository.AssignToLocation(ctx, input); err != nil {
		utils.LogError(err, "Failed to assign user to location")
		return err
	}

	return nil
}
