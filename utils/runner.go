package utils

import (
	"fmt"
	"sync"

	"github.com/pomcom/bagoScan/tools"
)

type Runner struct {
	Tools       []tools.Tool
	Filehandler Filehandler
}

func (r Runner) Run() {
	var wg sync.WaitGroup
	wg.Add(len(r.Tools))

	for _, t := range r.Tools {
		go func(tool tools.Tool) {
			defer wg.Done()
			output, err := tool.Execute("pomcom.digital")
			if err != nil {
				fmt.Println("Error in runner:", err)
				return
			}
			r.Filehandler.WriteToFile(tool.Name()+"-output.txt", output)
		}(t)
	}
	wg.Wait()
}
