package configreader

import (
	"os"

	"github.com/brkelkar/common_utils/logger"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

//Config provides option get Config variables from .yml file or
//from enviroment variable
//Incase of both eviroment variables overwrites .yml file variables
type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
		DBName   string `yaml:"database_name", envconfig:"DB_NAME"`
	} `yaml:"database"`
}

func processError(err error) {
	logger.Error("Error while reading config", err)
	os.Exit(2)
}

//ReadFile reads given config file
//and loads into config object
func (cfg *Config) ReadFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

//ReadEnv reads enviroment variables
//and loads into config object
func (cfg *Config) ReadEnv() {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
