package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime/trace"

	"github.com/pomcom/bagoScan/commands"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
	"github.com/pomcom/bagoScan/pkg/utils/monitoring"
)

func main() {

	// defer profile.Start().Stop()

	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	//
	// time.Sleep(5 * time.Second)

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
