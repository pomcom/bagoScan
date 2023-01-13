package main

import (
	"fmt"

	"github.com/pomcom/bagoScan/commands"
	"github.com/pomcom/bagoScan/pkg/utils"
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
