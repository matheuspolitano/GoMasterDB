package utils

import "github.com/spf13/viper"

type Config struct {
	DbDriver          string `mapstructure:"DB_DRIVER"`
	DbConnetionString string `mapstructure:"CON_STRING"`
}

func LoadConfig(path string) (conf Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		return
	}
	return

}
