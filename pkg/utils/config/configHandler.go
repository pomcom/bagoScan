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
	"io/ioutil"

	"github.com/pomcom/bagoScan/pkg/tools"
	"github.com/pomcom/bagoScan/pkg/tools/ffuf"
	"github.com/pomcom/bagoScan/pkg/tools/nmap"
	"github.com/pomcom/bagoScan/pkg/tools/nuclei"
	"github.com/pomcom/bagoScan/pkg/tools/sqlmap"
	"github.com/pomcom/bagoScan/pkg/tools/testssl"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
	"github.com/spf13/viper"
)

// Config containts the configuration of the tools to be run
type Config struct {
	ToolNames []string
	ToolMap   map[string]tools.Tool
	// AuthToken string `yaml:"auth_token"`
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
	"nuclei":  func(flags []string) tools.Tool { return nuclei.NewNuclei(flags, "testssl") },
	// "sqlmap":              func(flags []string) tools.Tool { return sqlmap.NewSQLMap(flags, "sqlmap") },
	"resource_discovery": func(flags []string) tools.Tool { return ffuf.NewResourceDiscovery(flags, "resource_discovery") },
	"ffufSqliApiTest":    func(flags []string) tools.Tool { return ffuf.NewSliApiTest(flags, "ffufSqliApiTest") },
	"sqlMapApiTest":      func(flags []string) tools.Tool { return sqlmap.NewSQLMapApiTest(flags, "sqlMapApiTest") },
}

// default tools that are executed when no config.yaml is provided
var defaultToolMap = map[string]tools.Tool{
	"nmap":    nmap.NewNmap(defaultToolFlags["nmap"].flags, defaultToolFlags["nmap"].name),
	"testssl": testssl.NewTestssl(defaultToolFlags["testssl"].flags, defaultToolFlags["testssl"].name),
	"nuclei":  nuclei.NewNuclei(defaultToolFlags["nuclei"].flags, defaultToolFlags["nuclei"].name),
	// "sqlmap":              sqlmap.NewSQLMap(defaultToolFlags["sqlmap"].flags, defaultToolFlags["sqlmap"].name),
	"resource_discovery": ffuf.NewResourceDiscovery(defaultToolFlags["resource_discovery"].flags, defaultToolFlags["resource_discovery"].name),
	"ffufSqliApiTest":    ffuf.NewSliApiTest(defaultToolFlags["ffufSqliApiTest"].flags, defaultToolFlags["ffufSqliApiTest"].name),
	"sqlMapApiTest":      sqlmap.NewSQLMapApiTest(defaultToolFlags["sqlMapApiTest"].flags, defaultToolFlags["sqlMapApiTest"].name),
}

// default flags that are used when no custom flags are provided in the config.yaml
var defaultToolFlags = map[string]struct {
	flags []string
	name  string
}{
	"nmap":    {[]string{"-T4", "-A"}, "nmap"},
	"testssl": {[]string{"--hints"}, "testssl"},
	"nuclei":  {[]string{"-u"}, "nuclei"},
	// "sqlmap":              {[]string{"-u"}, "sqlmap"},
	"resource_discovery": {[]string{"-w", "common.txt", "--recursion-depth", "3"}, "resource_discovery"},
	"ffufSqliApiTest":    {[]string{"-w", "payloads/sqli.txt"}, "ffufSqliApiTest"},
	"sqlMapApiTest":      {[]string{}, "sqlMapApiTest"},
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

	configHandler.readAuthToken()

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

func (configHandler ConfigHandler) readAuthToken() {
	if err := configHandler.viper.ReadInConfig(); err != nil {
		utils.Logger.Warn("Error reading auth token in config.yml")
	}
	authToken := configHandler.viper.GetString("auth_token")
	ioutil.WriteFile("output/auth_token.txt", []byte(authToken), 0644)
	utils.Logger.Info("auth token written to auth_token.txt")

}
