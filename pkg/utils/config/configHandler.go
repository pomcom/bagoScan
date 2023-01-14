package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Tools []string
}

type ConfigHandler struct {
	viper    *viper.Viper
	config   Config
	filepath string
}

func NewConfigHandler(filepath string) ConfigHandler {
	v := viper.New()
	v.SetConfigFile(filepath)
	v.ReadInConfig()
	return ConfigHandler{
		viper:    v,
		filepath: filepath,
	}
}

func (c ConfigHandler) ReadConfig() (Config, error) {
	c.viper.SetConfigFile(c.filepath)
	if err := c.viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	c.config = Config{
		Tools: c.viper.GetStringSlice("tools"),
	}
	return c.config, nil
}

func (c ConfigHandler) GetConfig() Config {
	return c.config
}
