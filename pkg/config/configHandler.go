package config

import (
	"github.com/spf13/viper"
)

type ConfigReader struct {
	configPath string
	viper      viper.Viper
}

func NewConfigReader(configPath string) ConfigReader {
	v := viper.New()
	v.SetConfigFile(configPath)
	return ConfigReader{
		configPath: configPath,
		viper:      *v,
	}
}

func (c ConfigReader) ReadConfig() error {
	if err := c.viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c ConfigReader) GetTools() []string {
	println("Get Tools did get called!!!")
	return c.viper.GetStringSlice("tools")

}
