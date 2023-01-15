package config

/*
Since I want to to use `ExecuteTool (tool, target)` in the testrunner
just with the name of the tool, a mapping needs to be done.

tool name <-> corresponding tool struct

Decided to do this in the config, since its a good way to seperate concerns.
*/

import (
	"github.com/pomcom/bagoScan/pkg/tools"
	"github.com/pomcom/bagoScan/pkg/tools/nmap"
	"github.com/pomcom/bagoScan/pkg/tools/testssl"
	"github.com/spf13/viper"
)

type Config struct {
	ToolNames []string
	ToolMap   map[string]tools.Tool
}

type ConfigHandler struct {
	viper    *viper.Viper
	config   Config
	filepath string
}

func NewConfigHandler(filepath string) *ConfigHandler {
	v := viper.New()
	v.SetConfigFile(filepath)
	v.ReadInConfig()
	return &ConfigHandler{
		viper:    v,
		filepath: filepath,
	}
}

func (c *ConfigHandler) ReadConfig() (Config, error) {
	c.viper.SetConfigFile(c.filepath)
	if err := c.viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	tools := c.viper.GetStringSlice("tools")
	toolMap := make(map[string]tools.Tool)
	for _, t := range tools {
		var tool tools.Tool
		switch t {
		case "testssl":
			tool = &testssl.Testssl{}
		case "nmap":
			tool = &nmap.Nmap{}
		}
		toolMap[t] = tool
	}
	c.config = Config{
		ToolNames: tools,
		ToolMap:   toolMap,
	}
	return c.config, nil
}
