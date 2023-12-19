package services

import (
	"context"
	"strconv"
	"time"

	"srating/domain"
	"srating/utils"
	"srating/x/rest"

	"github.com/hibiken/asynq"
)

type authService struct {
	userRepository domain.UserRepository
	asynqClient    *asynq.Client
	contextTimeout time.Duration
}

func NewAuthService(userRepository domain.UserRepository, asynqClient *asynq.Client, timeout time.Duration) domain.AuthService {
	return &authService{
		userRepository: userRepository,
		asynqClient:    asynqClient,
		contextTimeout: timeout,
	}
}

func (uu *authService) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	accessToken, err := utils.CreateAccessToken(user, secret, expiry)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (uu *authService) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	refreshToken, err := utils.CreateRefreshToken(user, secret, expiry)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (uu *authService) Register(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	if err := utils.Validate(user); err != nil {
		utils.LogError(err, "Failed to validate user")
		return err
	}
	// Hash the user's password
	hashedPassword, err := utils.GenerateHashPassword(user.Password)
	if err != nil {
		utils.LogError(err, "Failed to hash password")
		return err
	}
	// Set the hashed password
	user.Password = hashedPassword
	user.Role = domain.EmployeeRole
	user.Status = domain.ActiveStatus
	if err := uu.userRepository.UpdateUser(ctx, user); err != nil {
		utils.LogError(err, "Failed to update user")
		return err
	}

	// go func() {
	// 	code, err := utils.GenerateRandomString(8)
	// 	if err != nil {
	// 		utils.LogError(err, "Failed to generate random string")
	// 		return
	// 	}
	// 	if err := retry.Do(func() error {
	// 		taskSendVerifyEmailPayload := &tasks.TaskSendVerifyEmailPayload{
	// 			ID:       user.ID,
	// 			FullName: user.FullName,
	// 			Email:    user.Email,
	// 			Code:     code,
	// 		}
	// 		task, err := tasks.NewTaskSendVerifyEmail(taskSendVerifyEmailPayload)
	// 		if err != nil {
	// 			utils.LogError(err, "Failed to create task")
	// 			return err
	// 		}
	// 		info, err := uu.asynqClient.Enqueue(task)
	// 		if err != nil {
	// 			utils.LogError(err, "Failed to enqueue task")
	// 			return err
	// 		}
	// 		utils.LogData(info, "info task")
	// 		return nil
	// 	}); err != nil {
	// 		utils.LogError(err, "Failed to retry task send verify email")
	// 	}
	// }()

	// utils.LogInfo("Processed verify email task successfully")
	return nil
}

func (uu *authService) Login(c context.Context, input *domain.LoginRequest) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return nil, err
	}

	user, err := uu.userRepository.GetUserByUsername(ctx, input.Username)
	if err != nil {
		utils.LogError(err, "User not found")
		return nil, rest.ErrInvalidCredentials
	}
	// Verify the password
	if err := utils.CompareHashAndPassword(user.Password, input.Password); err != nil {
		utils.LogError(err, "Failed to verify password")
		return nil, rest.ErrInvalidCredentials
	}
	return user, nil
}

func (uu *authService) GetUserByID(c context.Context, id uint) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user, err := uu.userRepository.GetUserByID(ctx, id)
	if err != nil {
		utils.LogError(err, "Failed to retrieve user")
		return nil, err
	}
	return user, nil
}

func (uu *authService) ExtractIDFromToken(ctx context.Context, token, accessTokenSecret string) (uint, error) {
	rawID, err := utils.ExtractIDFromToken(token, accessTokenSecret)
	if err != nil {
		utils.LogError(err, "Failed to extract user ID from access token")
		return 0, err
	}
	if err != nil {
		utils.LogError(err, "Failed to extract user ID from refresh token")
		return 0, err
	}
	userID, err := strconv.Atoi(rawID)
	if err != nil {
		utils.LogError(err, "Failed to convert user ID to integer")
		return 0, err
	}
	return uint(userID), nil
}
