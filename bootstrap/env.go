package bootstrap

import (
	"srating/utils"

	"github.com/spf13/viper"
)

type Env struct {
	GinMode                 string `mapstructure:"GIN_MODE"`
	EmailSenderName         string `mapstructure:"EMAIL_SENDER_NAME"`
	BaseURL                 string `mapstructure:"BASE_URL"`
	DBHost                  string `mapstructure:"DB_HOST"`
	AppEnv                  string `mapstructure:"APP_ENV"`
	EmailSenderPassword     string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	DatabaseURL             string `mapstructure:"DB_URL"`
	ServerAddress           string `mapstructure:"SERVER_ADDRESS"`
	EmailSenderAddress      string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	DBUser                  string `mapstructure:"DB_USER"`
	AccessTokenSecret       string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret      string `mapstructure:"REFRESH_TOKEN_SECRET"`
	TimeZone                string `mapstructure:"TIME_ZONE"`
	RedisPassword           string `mapstructure:"REDIS_PASSWORD"`
	RedisURL                string `mapstructure:"REDIS_URL"`
	DBPassword              string `mapstructure:"DB_PASSWORD"`
	DBName                  string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour   int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	DBPort                  int    `mapstructure:"DB_PORT"`
	ReadTimeout             int    `mapstructure:"READ_TIMEOUT"`
	RequestTimeout          int    `mapstructure:"REQUEST_TIMEOUT"`
	RedisDB                 int    `mapstructure:"REDIS_DB"`
	RememberTokenExpiryHour int    `mapstructure:"REMEMBER_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour  int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	IdleTimeout             int    `mapstructure:"IDLE_TIMEOUT"`
	WriteTimeout            int    `mapstructure:"WRITE_TIMEOUT"`
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
