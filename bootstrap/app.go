package bootstrap

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	Env         *Env
	DB          *gorm.DB
	Redis       *redis.Client
	AsynqClient *asynq.Client
}

func NewApplication() *Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = NewPostgresDatabase(app.Env)
	// app.DB = NewMongoDatabase(app.Env)
	app.Redis = NewRedisClient(app.Env)
	app.AsynqClient = NewAsynqClient(app.Env)
	return app
}

func (a *Application) CloseConnection() {
	ClosePostgreConnection(a.DB)
	CloseRedisConnection(a.Redis)
	CloseAsynqConnection(a.AsynqClient)
}
