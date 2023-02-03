package services

import (
	"fmt"

	"github.com/pomcom/bagoScan/pkg/core"
	"github.com/pomcom/bagoScan/pkg/tools"
	"github.com/pomcom/bagoScan/pkg/utils/config"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
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

	// todo move this to own func.
	var outputFiles []string
	for _, output := range outputs {
		fileName := output.ToolName + "-" + output.Target + "-output.txt"

		err := service.fileHandler.WriteToFile(fileName, output.Result)
		utils.Logger.Info("result written to output file")
		if err != nil {
			return err
		}
		outputFiles = append(outputFiles, fileName)
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

	var outputFiles []string
	for _, output := range outputs {
		fileName := output.ToolName + "-" + output.Target + "-output.txt"
		err := service.fileHandler.WriteToFile(fileName, output.Result)
		utils.Logger.Info("result written to output file")
		if err != nil {
			return err
		}
		outputFiles = append(outputFiles, fileName)
		utils.Logger.Info("combines output file created")
	}

	err := service.fileHandler.CombineFiles(outputFiles, "output/combined-output.txt")
	if err != nil {
		return err
	}

	return nil

}

func NewTestRunnerService(configPath string) TestRunnerService {
	configHandler := config.NewConfigHandler(configPath)
	config := configHandler.ReadConfig()
	fileHandler := core.NewFilehandler("")
	return TestRunnerService{configHandler: configHandler, fileHandler: fileHandler, config: config}
}
