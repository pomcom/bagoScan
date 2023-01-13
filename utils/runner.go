package utils

import (
	"fmt"
	"sync"

	"github.com/pomcom/bagoScan/tools"
	"github.com/pomcom/bagoScan/tools/nmap"
	"github.com/pomcom/bagoScan/tools/testssl"
)

type Runner struct {
	Tools       []tools.Tool
	Filehandler Filehandler
}

func (r Runner) Run(target string) {
	var wg sync.WaitGroup
	wg.Add(len(r.Tools))

	for _, t := range r.Tools {
		go func(tool tools.Tool) {
			defer wg.Done()
			output, err := tool.Execute(target)
			if err != nil {
				fmt.Println("Error in runner:", err)
				return
			}
			r.Filehandler.WriteToFile(tool.Name()+"-output.txt", output)
		}(t)
	}
	wg.Wait()
}

func NewRunner(tools []tools.Tool, filehandler Filehandler) *Runner {
	return &Runner{
		Tools:       tools,
		Filehandler: filehandler,
	}
}

var runner *Runner

func Init() {
	runner = NewRunner([]tools.Tool{
		testssl.Testssl{},
		nmap.Nmap{},
	}, Filehandler{})
}
