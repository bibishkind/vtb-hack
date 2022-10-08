package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type (
	Config struct {
		Postgres       *PostgresConfig
		HTTP           *HTTPConfig
		RequestTimeout time.Duration
	}

	PostgresConfig struct {
		URI string
	}

	HTTPConfig struct {
		Port           string
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		MaxHeaderBytes int
	}
)

func Init(configPath string) (*Config, error) {
	path := strings.Split(configPath, "/")

	viper.AddConfigPath(strings.Join(path[0:len(path)-1], "/"))
	viper.SetConfigName(path[len(path)-1])

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{
		Postgres: &PostgresConfig{
			URI: os.Getenv("DB_URL"),
		},
		HTTP: &HTTPConfig{
			Port:           viper.GetString("http.port"),
			ReadTimeout:    viper.GetDuration("http.readTimeout"),
			WriteTimeout:   viper.GetDuration("http.writeTimeout"),
			MaxHeaderBytes: viper.GetInt("http.maxHeaderBytes"),
		},
		RequestTimeout: viper.GetDuration("requestTimeout"),
	}

	return cfg, nil
}
