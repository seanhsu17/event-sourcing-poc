package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	fileType = "yaml"
)

var (
	C = new(Config)
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	BaseRoute     string
	Port          int
	RunMode       string
	Name          string
	EnableProfile bool
}

func InitConfigs() {
	fromFile()
	fmt.Printf("C.App = %+v\n", C.App)
}

func fromFile() {
	viper.SetConfigType(fileType)
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Read configs error，reason：" + err.Error())
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		panic("Unmarshal error: " + err.Error())
	}
}
