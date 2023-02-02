package config

/*
Since I want to to use `ExecuteTool (tool, target)` in the testrunner
just with the name of the tool, a mapping needs to be done.

tool name <-> corresponding tool struct

Decided to do this in the config, since its a good way to seperate concerns.
maybe change struct literal syntax to NewTestssl() function, that needs to be implemented

%TODO rework maybe..

*/

import (
	"github.com/pomcom/bagoScan/pkg/tools"
	"github.com/pomcom/bagoScan/pkg/tools/nmap"
	"github.com/pomcom/bagoScan/pkg/tools/nuclei"
	"github.com/pomcom/bagoScan/pkg/tools/testssl"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
	"github.com/spf13/viper"
)

// Config containts the configuration of the tools to be run
type Config struct {
	ToolNames []string
	ToolMap   map[string]tools.Tool
}

// ConfigHandler reads and parses the config file
type ConfigHandler struct {
	viper    *viper.Viper
	config   Config
	filepath string
}

// all implemented tools need to be initialized and mapped here
var defaultToolFactories = map[string]func([]string) tools.Tool{
	"nmap":    func(flags []string) tools.Tool { return nmap.NewNmap(flags, "nmap") },
	"testssl": func(flags []string) tools.Tool { return testssl.NewTestssl(flags, "testssl") },
	"nuclei":  func(flags []string) tools.Tool { return testssl.NewTestssl(flags, "testssl") },
}

// default tools that are executed when no config.yaml is provided
var defaultToolMap = map[string]tools.Tool{
	"nmap":    nmap.NewNmap(defaultToolFlags["nmap"].flags, defaultToolFlags["nmap"].name),
	"testssl": testssl.NewTestssl(defaultToolFlags["testssl"].flags, defaultToolFlags["testssl"].name),
	"nuclei":  nuclei.NewNuclei(defaultToolFlags["nuclei"].flags, defaultToolFlags["nuclei"].name),
}

// default flags that are used when no custom flags are provided in the config.yaml
var defaultToolFlags = map[string]struct {
	flags []string
	name  string
}{
	"nmap":    {[]string{"-T4", "-A"}, "nmap"},
	"testssl": {[]string{"--hints"}, "testssl"},
}

func NewConfigHandler(filepath string) ConfigHandler {
	v := viper.New()
	v.SetConfigFile(filepath)

	if err := v.ReadInConfig(); err != nil {
		utils.Logger.Info("No config file provided - using default settings")
	} else {
		utils.Logger.Info("Using provided configuration file")
	}

	return ConfigHandler{
		viper:    v,
		filepath: filepath,
	}
}

// Returns an empty struct if no file is provided
func (configHandler ConfigHandler) ReadConfig() Config {

	// checking for error to use own logger and not vipers build in logging
	if err := configHandler.viper.ReadInConfig(); err != nil {
		return Config{ToolMap: defaultToolMap}
	}

	toolNames := configHandler.viper.GetStringSlice("tools")
	toolFactories := defaultToolFactories

	toolFlags := make(map[string][]string)
	toolMap := make(map[string]tools.Tool)

	for _, t := range toolNames {
		factory, ok := toolFactories[t]
		if !ok {
			utils.Logger.Warn("Tool not found - Check config.yml for typos. Has a new tool been implemented and added in the configHandler?")
			return Config{}
		}
		toolFlags[t] = configHandler.viper.GetStringSlice(t)
		tool := factory(toolFlags[t])

		toolMap[t] = tool
	}
	configHandler.config = Config{
		ToolNames: toolNames,
		ToolMap:   toolMap,
	}
	return configHandler.config
}
