package main

import (
	"github.com/pomcom/bagoScan/tools"
	"github.com/pomcom/bagoScan/tools/testssl"
	"github.com/pomcom/bagoScan/utils"
)

func main() {
	var tools []tools.Tool
	fileHandler := utils.Filehandler{}

	tools = append(tools, testssl.Testssl{})

	r := utils.Runner{Tools: tools, Filehandler: fileHandler}
	r.Run()
}
