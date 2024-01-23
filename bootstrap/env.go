package bootstrap

import (
	"srating/utils"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                  string `mapstructure:"APP_ENV"`
	ServerAddress           string `mapstructure:"SERVER_ADDRESS"`
	RequestTimeout          int    `mapstructure:"REQUEST_TIMEOUT"`
	ReadTimeout             int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout            int    `mapstructure:"WRITE_TIMEOUT"`
	IdleTimeout             int    `mapstructure:"IDLE_TIMEOUT"`
	DatabaseURL             string `mapstructure:"DB_URL"`
	AccessTokenExpiryHour   int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour  int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	RememberTokenExpiryHour int    `mapstructure:"REMEMBER_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret       string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret      string `mapstructure:"REFRESH_TOKEN_SECRET"`
	GinMode                 string `mapstructure:"GIN_MODE"`
	BaseURL                 string `mapstructure:"BASE_URL"`
	DBHost                  string `mapstructure:"DB_HOST"`
	DBPort                  int    `mapstructure:"DB_PORT"`
	DBName                  string `mapstructure:"DB_NAME"`
	DBUser                  string `mapstructure:"DB_USER"`
	DBPassword              string `mapstructure:"DB_PASSWORD"`
	RedisURL                string `mapstructure:"REDIS_URL"`
	RedisPassword           string `mapstructure:"REDIS_PASSWORD"`
	RedisDB                 int    `mapstructure:"REDIS_DB"`
	EmailSenderName         string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress      string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword     string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	TimeZone                string `mapstructure:"TIME_ZONE"`
}

// InitViper initializes the Viper configuration.

// NewEnv creates and returns a new Env object.
func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		utils.LogFatal(err, "Error reading config file")
	}

	if err := viper.Unmarshal(&env); err != nil {
		utils.LogFatal(err, "Failed to unmarshal config file")
	}
	// Add any additional initialization or validation logic here
	return &env
}
