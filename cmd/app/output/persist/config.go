package persist

import (
	"TgBot/cmd/core/constant"
	"github.com/spf13/viper"
	"log"
)

type DatabaseConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	DbName    string `mapstructure:"db_name"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Schema    string `mapstructure:"schema"`
	BatchSize int    `mapstructure:"batch-size"`
}

func InitConfig() *DatabaseConfig {
	config := DatabaseConfig{}
	viper.SetConfigFile(constant.ConfigFilePath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't load config file", err)
	}
	err = viper.UnmarshalKey("database", &config)
	if err != nil {
		log.Fatal("Can't unmarshal config file", err)
	}
	return &config
}
