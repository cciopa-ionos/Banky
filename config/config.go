package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName(".banky")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // current directory

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
