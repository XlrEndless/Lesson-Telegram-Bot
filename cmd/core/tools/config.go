package tools

import (
	"TgBot/cmd/core/constant"
	"github.com/spf13/viper"
	"log"
)

type CommonConfig struct {
	DefaultDatePageLimit int   `mapstructure:"default-date-page-limit"`
	DefaultTimePageLimit int   `mapstructure:"default-time-page-limit"`
	PollIntervalMillis   int64 `mapstructure:"poll-interval-millis"`
}

func InitConfig() *CommonConfig {
	config := CommonConfig{}
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
