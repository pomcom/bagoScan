package config

import (
	"github.com/spf13/viper"
)

type ConfigHandler struct {
	viper *viper.Viper
}

func NewConfigHandler(configPath string) ConfigHandler {
	myViper := viper.New()
	myViper.SetConfigFile(configPath)
	myViper.ReadInConfig()
	return ConfigHandler{
		viper: myViper,
	}
}

func (c ConfigHandler) GetTools() []string {
	return c.viper.GetStringSlice("tools")
}
