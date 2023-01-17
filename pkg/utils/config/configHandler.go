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
	"github.com/pomcom/bagoScan/pkg/tools/testssl"
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

// all implemented tools need to be initialized and mapped here
var defaultToolFactories = map[string]func([]string) tools.Tool{
	"nmap":    func(flags []string) tools.Tool { return nmap.NewNmap(flags, "nmap") },
	"testssl": func(flags []string) tools.Tool { return testssl.NewTestssl(flags, "testssl") },
}

// default tools that are executed when no config.yaml is provided
var defaultToolMap = map[string]tools.Tool{
	"nmap":    nmap.NewNmap(defaultToolFlags["nmap"].flags, defaultToolFlags["nmap"].name),
	"testssl": testssl.NewTestssl(defaultToolFlags["testssl"].flags, defaultToolFlags["testssl"].name),
}

// default flags that are used when no custom flags are provided in the config.yaml
var defaultToolFlags = map[string]struct {
	flags []string
	name  string
}{
	"nmap":    {[]string{"-T4", "-A"}, "nmap"},
	"testssl": {[]string{"--json"}, "testssl"},
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
		// TODO
		// utils.Logger.Info("No config file provided - using default tools")
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
			utils.Logger.Warn("Tool not found - Check config.yml for typos. Has a new tool been implemented and added in the configHandler?")
			return Config{}, fmt.Errorf("tool not found: %s", t)
		}
		toolFlags[t] = configHandler.viper.GetStringSlice(t)
		tool := factory(toolFlags[t])

		toolMap[t] = tool
	}
	configHandler.config = Config{
		ToolNames: toolNames,
		ToolMap:   toolMap,
	}
	return configHandler.config, nil
}

func (configHandler ConfigHandler) ReadSingleToolConfig(toolName string) (tools.Tool, error) {
	defaultFlag, ok := defaultToolFlags[toolName]
	if !ok {
		return nil, fmt.Errorf("tool not supported: %s", toolName)
	}
	factory, ok := defaultToolFactories[toolName]
	if !ok {
		return nil, fmt.Errorf("tool not found: %s", toolName)
	}
	tool := factory(defaultFlag.flags)
	return tool, nil
}
