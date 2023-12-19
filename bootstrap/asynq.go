package bootstrap

import (
	"io"

	"srating/utils"

	"github.com/hibiken/asynq"
)

func NewAsynqClient(env *Env) *asynq.Client {
	opt := asynq.RedisClientOpt{Addr: env.RedisURL, Password: env.RedisPassword, DB: env.RedisDB}
	asynqClient := asynq.NewClient(opt)
	return asynqClient
}

func CloseAsynqConnection(asynqClient io.Closer) {
	err := asynqClient.Close()
	if err != nil {
		utils.LogFatal(err, "Failed to close Asynq connection")
		return
	}
}
