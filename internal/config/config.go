package config

import "github.com/spf13/viper"

func LoadViperEnv() {
	viper.AddConfigPath("/")
	viper.SetConfigFile("env.yaml")
	viper.AutomaticEnv()

	envException := viper.ReadInConfig()
	
	if envException != nil {
		panic(envException)
	}
}