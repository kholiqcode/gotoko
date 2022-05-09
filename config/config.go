package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

var cfg Config
var doOnce sync.Once

type Config struct {
	AppEnv  string `mapstructure:"APP_ENV,default=development"`
	AppHost string `mapstructure:"APP_HOST,default=127.0.0.1"`
	AppPort int    `mapstructure:"APP_PORT,default=8080"`

	DBDriver   string `mapstructure:"DB_DRIVER,default=mysql"`
	DBHost     string `mapstructure:"DB_HOST,default=127.0.0.1"`
	DBPort     string `mapstructure:"DB_PORT,default=3306"`
	DBName     string `mapstructure:"DB_NAME,default=alta_code_structuring"`
	DBUser     string `mapstructure:"DB_USER,default=root"`
	DBPassword string `mapstructure:"DB_PASSWORD,default="`

	RedisHost     string `mapstructure:"REDIS_HOST,default=127.0.0.1"`
	RedisPort     string `mapstructure:"REDIS_PORT,default=6379"`
	RedisPassword string `mapstructure:"REDIS_Password,default="`

	TestDBDriver   string `mapstructure:"TEST_DB_DRIVER,default=mysql"`
	TestDBHost     string `mapstructure:"TEST_DB_HOST,default=127.0.0.1"`
	TestDBPort     string `mapstructure:"TEST_DB_PORT,default=3306"`
	TestDBName     string `mapstructure:"TEST_DB_NAME,default=alta_code_structuring_test"`
	TestDBUser     string `mapstructure:"TEST_DB_USER,default=root"`
	TestDBPassword string `mapstructure:"TEST_DB_PASSWORD,default="`

	JwtAccessExpired  string `mapstructure:"JWT_ACCESS_EXPIRED",env:"JWT_ACCESS_EXPIRED,default=3600"`
	JwtRefreshExpired string `mapstructure:"JWT_REFRESH_EXPIRED",env:"JWT_REFRESH_EXPIRED,default=3600"`
	JwtSecretKey      string `mapstructure:"JWT_SECRET_KEY",env:"JWT_SECRET_KEY,default=bismillah"`

	DeployedAt string
}

func NewConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Err(err).Msgf("cannot read .env file: %v", err)
	}

	doOnce.Do(func() {
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Err(err).Msg("cannot unmarshaling config")
		}
	})

	cfg.AppHost = mustHostname()

	cfg.DeployedAt = time.Now().String()

	return &cfg
}

func Get() Config {
	return cfg
}

func mustHostname() string {
	host, err := os.Hostname()
	if err != nil {
		log.Err(err).Msgf("hostnameError: %v", err)
	}
	return host
}
