package services

import (
	"github.com/pomcom/bagoScan/pkg/core"
	"github.com/pomcom/bagoScan/pkg/tools"
	"github.com/pomcom/bagoScan/pkg/utils"
	"github.com/pomcom/bagoScan/pkg/utils/config"
	"github.com/spf13/viper"
)

type TestRunnerService struct {
	configHandler config.ConfigHandler
	runner        core.TestRunner
	fileHandler   core.Filehandler
}

func (service TestRunnerService) RunAllTools(target string) error {
	config, err := service.configHandler.ReadConfig()
	if err != nil {
		return err
	}
	service.runner = core.NewTestRunner(config.ToolMap)
	outputs := service.runner.Run(target)

	for _, output := range outputs {
		fileName := output.ToolName + "-" + target + "-output.txt"
		err := service.fileHandler.WriteToFile(fileName, output.Result)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service TestRunnerService) RunSingleTool(toolName string, target string) error {
	config, err := service.configHandler.ReadConfig()
	if err != nil {
		// check if file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			tool, err := service.configHandler.ReadSingleToolConfig(toolName)
			if err != nil {
				return err
			}
			singleToolMap := map[string]tools.Tool{toolName: tool}
			runner := utils.NewTestRunner(singleToolMap)
			outputs := runner.Run(target)
			for _, output := range outputs {
				fileName := output.ToolName + "-" + target + "-output.txt"
				err := service.fileHandler.WriteToFile(fileName, output.Result)
				if err != nil {
					return err
				}
			}
			return nil
		}
		return err
	}
	// create a new TestRunner with only the specified tool
	singleToolMap := map[string]tools.Tool{toolName: config.ToolMap[toolName]}
	runner := utils.NewTestRunner(singleToolMap)
	outputs := runner.Run(target)
	for _, output := range outputs {
		fileName := output.ToolName + "-" + target + "-output.txt"
		err := service.fileHandler.WriteToFile(fileName, output.Result)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewTestRunnerService(configPath string) TestRunnerService {
	configHandler := config.NewConfigHandler(configPath)
	config, err := configHandler.ReadConfig()
	if err != nil {
		panic(err)
	}
	fileHandler := core.NewFilehandler("")
	return TestRunnerService{configHandler: configHandler, fileHandler: fileHandler, runner: core.NewTestRunner(config.ToolMap)}
}
