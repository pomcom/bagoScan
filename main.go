package main

import (
	"fmt"

	"github.com/pomcom/bagoScan/commands"
	"github.com/pomcom/bagoScan/pkg/utils/config"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type PentestService struct {
	configHandler *config.ConfigHandler
}

func (s *PentestService) RunScan(target string) {
	config, err := s.configHandler.ReadConfig()
	tools := config.Tools
	runner := utils.NewTestRunner(tools)
	outputs := runner.Run(target)

	for _, output := range outputs {
		fmt.Println(output.ToolName + ": " + output.Result)
	}
}

func NewPentestService(configPath string) *PentestService {
	configHandler := config.NewConfigHandler(configPath)
	return &PentestService{configHandler: &configHandler}
}

func main() {

	utils.InitializeLogger()

	utils.Logger.Info("Logger initalized")
	utils.Logger.Info("Application started")

	err := commands.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}

}
