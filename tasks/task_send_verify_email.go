package tasks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"srating/domain"
	"srating/mail"
	"srating/utils"

	"github.com/hibiken/asynq"
)

const (
	TaskSendVerifyEmail = "task:verify_email"
)

type TaskSendVerifyEmailPayload struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type ProcessTaskSendVerifyEmail struct {
	userService domain.UserService
	mailer      mail.EmailSender
}

func NewProcessTaskSendVerifyEmail(userService domain.UserService, mailer mail.EmailSender) *ProcessTaskSendVerifyEmail {
	return &ProcessTaskSendVerifyEmail{
		userService: userService,
		mailer:      mailer,
	}
}

func NewTaskSendVerifyEmail(payload *TaskSendVerifyEmailPayload) (*asynq.Task, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(&payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TaskSendVerifyEmail, b.Bytes()), nil
}

func (processor *ProcessTaskSendVerifyEmail) ProcessTask(ctx context.Context, task *asynq.Task) error {
	utils.LogInfo("Processing verify email task")
	var payload TaskSendVerifyEmailPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		utils.LogError(err, "Error unmarshaling payload")
		return err
	}

	user, err := processor.userService.GetUserByID(ctx, payload.ID)
	if err != nil {
		utils.LogError(err, "Error getting user by ID")
		return err
	}

	subject := "Welcome to Tomotek"
	content := fmt.Sprintf(`Hello %s, <br>
		Welcome to Tomotek. Please verify your email by clicking the link below: <br>
		<a href="%s">Verify</a> <br>`, payload.FullName, payload.Code)
	to := []string{user.Email}

	err = processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		utils.LogError(err, "Error sending verify email")
		return err
	}

	utils.LogInfo("Processed verify email task successfully")
	return nil
}
