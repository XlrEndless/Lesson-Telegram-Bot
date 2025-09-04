package rest

import (
	"TgBot/cmd/core/constant"
	"github.com/spf13/viper"
	"log"
)

type TgConfig struct {
	BotToken string `mapstructure:"bot-token"`
	BasePath string `mapstructure:"base-path"`
}

func InitConfig() *TgConfig {
	config := TgConfig{}
	viper.SetConfigFile(constant.ConfigFilePath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't load config file", err)
	}
	err = viper.UnmarshalKey("tg", &config)
	if err != nil {
		log.Fatal("Can't unmarshal config file", err)
	}
	return &config
}
