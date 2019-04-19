package config

import (
	"encoding/json"
	"fmt"
	"os"
)

/*type Config struct {
	DB *DBConfig
}*/

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		Dialect  string `json:"dialect"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Charset  string `json:"charset"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	error := jsonParser.Decode(&config)
	return config, error
}

func GetConfig() Config {
	config, _ := LoadConfiguration("api/config/config.json")
	return config
}
