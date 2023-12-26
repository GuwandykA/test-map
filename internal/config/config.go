package config

import (
	"bd-backend/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"localhost"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	PublicFilePath       string        `yaml:"public_file_path" env-required:"true"`
	LogPath              string        `yaml:"log_path" env-required:"true"`
	Storage              StorageConfig `yaml:"storage"`
	JwtKeySupAdmin       string        `yaml:"jwt_key_sup_admin" env-required:"true"`
	JwtKey               string        `yaml:"jwt_key" env-required:"true"`
	JwtKeyForgetPassword string        `yaml:"jwt_key_forget_password" env-required:"true"`
	AppVersion           string        `yaml:"app_version" env-required:"true"`
}

type StorageConfig struct {
	PgPoolMaxConn int    `yaml:"pg_pool_max_conn" env-required:"true"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Database      string `json:"database"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		pathConfig := "./config.yml"
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig(pathConfig, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
