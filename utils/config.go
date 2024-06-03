package utils

import "github.com/spf13/viper"

const configFile = ".env"

type Config struct {
	AccessTokenSecret       string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RedisPassword           string `mapstructure:"REDIS_PASSWORD"`
	RefreshTokenSecret      string `mapstructure:"REFRESH_TOKEN_SECRET"`
	GinMode                 string `mapstructure:"GIN_MODE"`
	DBUrl                   string `mapstructure:"DB_URL"`
	AppEnv                  string `mapstructure:"APP_ENV"`
	EmailSenderAddress      string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderName         string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderPassword     string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	ServerAddress           string `mapstructure:"SERVER_ADDRESS"`
	Host                    string `mapstructure:"HOST"`
	BaseURL                 string `mapstructure:"BASE_URL"`
	RedisURL                string `mapstructure:"REDIS_URL"`
	ContextTimeout          int    `mapstructure:"CONTEXT_TIMEOUT"`
	RedisDB                 int    `mapstructure:"REDIS_DB"`
	RememberTokenExpiryHour int    `mapstructure:"REMEMBER_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour  int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenExpiryHour   int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
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
