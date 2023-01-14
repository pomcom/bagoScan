package main

import (
	"fmt"

	"github.com/pomcom/bagoScan/commands"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

func main() {

	utils.InitializeLogger()

	utils.Logger.Info("Logger initalized")
	utils.Logger.Info("Application started")

	err := commands.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}

}
