// ready

package config

import (
	"citizenship/pkg/logging"
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type SitesToParse struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	TempPath string `yaml:"temppath"`
	Path     string `yaml:"path"`
}

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
	Services []SitesToParse `yaml:"sitestoparse"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("reading application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			logger.Fatal(fmt.Sprintf("failed to read config: %v", err))
		}
	})
	return instance
}
