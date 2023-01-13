package main

import (
	"fmt"

	"github.com/pomcom/bagoScan/commands"
)

func main() {
	// var tools []tools.Tool
	// fileHandler := utils.Filehandler{}
	//
	// tools = append(tools, testssl.Testssl{})
	//
	// runner := utils.Runner{Tools: tools, Filehandler: fileHandler}
	// runner.Run()

	err := commands.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}

}
