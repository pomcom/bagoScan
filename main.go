package main

import (
	"fmt"

	"github.com/pomcom/bagoScan/commands"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	cmdError := commands.Execute()
	if cmdError != nil && cmdError.Error() != "" {
		fmt.Println(cmdError)
	}

}
