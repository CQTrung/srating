package utils

import "github.com/spf13/viper"

const configFile = ".env"

type Config struct {
	AppEnv                  string `mapstructure:"APP_ENV"`
	Host                    string `mapstructure:"HOST"`
	ServerAddress           string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout          int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBUrl                   string `mapstructure:"DB_URL"`
	AccessTokenExpiryHour   int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour  int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	RememberTokenExpiryHour int    `mapstructure:"REMEMBER_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret       string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret      string `mapstructure:"REFRESH_TOKEN_SECRET"`
	GinMode                 string `mapstructure:"GIN_MODE"`
	BaseURL                 string `mapstructure:"BASE_URL"`
	RedisURL                string `mapstructure:"REDIS_URL"`
	RedisPassword           string `mapstructure:"REDIS_PASSWORD"`
	RedisDB                 int    `mapstructure:"REDIS_DB"`
	EmailSenderName         string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress      string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword     string `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

// NewEnv creates and returns a new Env object.
func LoadConfig(path string) (err error) {
	var config Config
	viper.SetConfigFile(configFile)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		LogFatal(err, "Failed to read config file")
	}
	if err := viper.Unmarshal(&config); err != nil {
		LogFatal(err, "Failed to unmarshal config file")
	}
	// Add any additional initialization or validation logic here
	return nil
}
