package bootstrap

import (
	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func NewApplication() *Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = NewPostgresDatabase(app.Env)
	return app
}

func (a *Application) CloseConnection() {
	ClosePostgreConnection(a.DB)
}
