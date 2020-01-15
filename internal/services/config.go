package services

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const configFilePath string = "config/config.yaml"

var Config ConfigFile

type ConfigFile struct {
	DBName         string `yaml:"db_name"`
	DataSourceName string `yaml:"db_data_source_name"`
	JWTKey         string `yaml:"jwt_key"`
}

func InitConfig() {
	config, _ := ioutil.ReadFile(configFilePath)

	var configFile ConfigFile
	yaml.Unmarshal(config, &configFile)

	Config = configFile
}
