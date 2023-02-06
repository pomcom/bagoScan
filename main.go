package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime/trace"

	"github.com/pomcom/bagoScan/pkg/commands"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
	"github.com/pomcom/bagoScan/pkg/utils/monitoring"
)

func main() {

	file, err := os.Create("trace.out")
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not create trace output file: ", err)
	}
	defer file.Close()
	trace.Start(file)
	defer trace.Stop()

	monitoring.StartPrometheusServer()

	utils.InitializeLogger()

	utils.Logger.Info("Logger initalized")
	utils.Logger.Info("Application started")

	errCmd := commands.Execute()
	if errCmd != nil && err.Error() != "" {
		fmt.Println(errCmd)
	}
}
