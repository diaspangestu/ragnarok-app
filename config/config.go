package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string `env:"APP_NAME"`
	Port       string `env:"PORT"`
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
	DbTimezone string `env:"DB_TIMEZONE"`
}

var Cfg *Config

func LoadConfig() (*Config, error) {
	//envPath := os.Getenv("ENV_PATH")
	//if envPath == "" {
	//	envPath = ".env"
	//}
	//
	//viper.SetConfigFile(envPath)
	//viper.SetConfigType("env")
	//viper.AddConfigPath(".")
	//viper.AutomaticEnv()
	//
	//viper.SetConfigFile(".env")
	//viper.SetConfigType("env")
	//viper.AddConfigPath(".")
	//viper.AutomaticEnv()
	//
	//if err := viper.ReadInConfig(); err != nil {
	//	return nil, fmt.Errorf("failed to read config: %w", err)
	//}
	//
	//var cfg Config
	//if err := viper.Unmarshal(&cfg); err != nil {
	//	return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	//}

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err.Error()))
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("Error parsing env: %s", err.Error()))
	}

	Cfg = &cfg
	return &cfg, nil
}
