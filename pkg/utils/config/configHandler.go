package config

/*
Since I want to to use `ExecuteTool (tool, target)` in the testrunner
just with the name of the tool, a mapping needs to be done.

tool name <-> corresponding tool struct

Decided to do this in the config, since its a good way to seperate concerns.
maybe change struct literal syntax to NewTestssl() function, that needs to be implemented

*/

import (
	"fmt"

	"github.com/pomcom/bagoScan/pkg/tools"
	"github.com/pomcom/bagoScan/pkg/tools/nmap"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
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

var defaultToolMap = map[string]tools.Tool{
	"nmap": &nmap.Nmap{},
	// "testssl": testssl.Testssl{},
}

// all implemented tools need to be initialized here
var defaultToolFactories = map[string]func() tools.Tool{
	// "testssl": func() tools.Tool { return testssl.Testssl{} },
	"nmap": func() tools.Tool { return &nmap.Nmap{} },
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

func (configHandler ConfigHandler) ReadConfig() (Config, error) {

	configHandler.viper.SetConfigFile(configHandler.filepath)
	if err := configHandler.viper.ReadInConfig(); err != nil {
		utils.Logger.Info("No config file provided - using default tools")
		return Config{ToolMap: defaultToolMap}, nil
	}

	toolNames := configHandler.viper.GetStringSlice("tools")
	utils.Logger.Info("Using provided configuration file")
	toolFactories := defaultToolFactories

	toolFlags := make(map[string][]string)
	toolMap := make(map[string]tools.Tool)

	for _, t := range toolNames {
		factory, ok := toolFactories[t]
		if !ok {
			utils.Logger.Warn("Tool not found - Has a new tool been implemented and added in the configHandler?")
			return Config{}, fmt.Errorf("tool not found: %s", t)
		}
		toolFlags[t] = configHandler.viper.GetStringSlice(t)
		tool := factory()

		tool.SetFlags(toolFlags[t])
		fmt.Println("tool flags in configHandler:", toolFlags)
		toolMap[t] = tool
	}
	configHandler.config = Config{
		ToolNames: toolNames,
		ToolMap:   toolMap,
	}
	return configHandler.config, nil
}
