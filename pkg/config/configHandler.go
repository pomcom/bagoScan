package config

import (
	"github.com/spf13/viper"
)

type ConfigHandler struct {
	viper *viper.Viper
}

func NewConfigHandler(configPath string) *ConfigHandler {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.ReadInConfig()
	return &ConfigHandler{
		viper: v,
	}
}

func (c *ConfigHandler) GetTools() []string {
	return c.viper.GetStringSlice("tools")
}
