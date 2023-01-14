package utils

import (
	"fmt"
	"sync"
)

type TestRunner struct {
	Tools       []string
	Filehandler Filehandler
}

func (r TestRunner) Run(target string) {
	var wg sync.WaitGroup
	wg.Add(len(r.Tools))

	for _, t := range r.Tools {
		go func(tool string) {
			defer wg.Done()
			output, err := ExecuteTool(tool, target)
			if err != nil {
				fmt.Println("Error in runner:", err)
				return
			}
			r.Filehandler.WriteToFile(tool+"-output.txt", output)
		}(t)
	}
	wg.Wait()
}

func NewTestRunner(tools []string) *TestRunner {
	return &TestRunner{
		Tools:       tools,
		Filehandler: Filehandler{},
	}
}
