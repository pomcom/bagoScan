package utils

import (
	"fmt"
	"sync"

	"github.com/pomcom/bagoScan/pkg/tools"
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

func NewRunner() *Runner {
	return &Runner{
		Tools:       make([]tools.Tool, 0),
		Filehandler: Filehandler{},
	}
}

func (r *Runner) AddTool(tools ...tools.Tool) {
	r.Tools = append(r.Tools, tools...)
}
