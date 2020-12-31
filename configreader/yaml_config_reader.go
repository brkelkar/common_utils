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
		Port int    `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`

	DatabaseName struct {
		AwacsDBName         string `yaml:"acawsDBName", envconfig:"AWACS_DB"`
		AwacsSmartDBName    string `yaml:"acawsSmartDBName", envconfig:"AWACS_SMART_DB"`
		SmartStockistDBName string `yaml:"smartStockistDBName", envconfig:"AWACS_SMART_STOCKIST_DB"`
	} `yaml:"databaseName"`

	DatabaseConfig struct {
		Port     int    `yaml:"port", envconfig:"DB_PORT"`
		Host     string `yaml:"host", envconfig:"DB_HOST"`
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
	} `yaml:"databaseConfig"`
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
