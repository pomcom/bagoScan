package services

import (
	"fmt"

	"github.com/pomcom/bagoScan/pkg/core"
	"github.com/pomcom/bagoScan/pkg/tools"
	"github.com/pomcom/bagoScan/pkg/utils/config"
)

type TestRunnerService struct {
	configHandler config.ConfigHandler
	config        config.Config
	runner        core.TestRunner
	fileHandler   core.Filehandler
}

// no error handling for ReadConfig() needed, since it returns empty valid struct in err case
func (service TestRunnerService) RunAllTools(targets []string) error {
	runner := core.NewTestRunner(service.config.ToolMap)
	outputs := runner.Run(targets)

	for _, output := range outputs {
		fileName := output.ToolName + "-" + output.Target + "-output.txt"
		err := service.fileHandler.WriteToFile(fileName, output.Result)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service TestRunnerService) RunSingleTool(toolName string, targets []string) error {
	tool, ok := service.config.ToolMap[toolName]
	if !ok {
		return fmt.Errorf("tool not found: %s", toolName)
	}

	singleToolMap := map[string]tools.Tool{toolName: tool}
	runner := core.NewTestRunner(singleToolMap)
	outputs := runner.Run(targets)
	println("hi")

	for _, output := range outputs {
		fileName := output.ToolName + "-" + output.Target + "-output.txt"
		err := service.fileHandler.WriteToFile(fileName, output.Result)
		if err != nil {
			return err
		}
	}
	return nil

}

func NewTestRunnerService(configPath string) TestRunnerService {
	configHandler := config.NewConfigHandler(configPath)
	config := configHandler.ReadConfig()
	fileHandler := core.NewFilehandler("")
	return TestRunnerService{configHandler: configHandler, fileHandler: fileHandler, config: config}
}
