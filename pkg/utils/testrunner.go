package utils

import (
	"fmt"
	"sync"
)

type Output struct {
	ToolName string
	Result   string
}

type TestRunner struct {
	Tools []string
}

func (r TestRunner) Run(target string) []Output {
	var wg sync.WaitGroup
	wg.Add(len(r.Tools))
	outputs := make([]Output, 0)

	for _, t := range r.Tools {
		go func(tool string) {
			defer wg.Done()
			result, err := ExecuteTool(tool, target)
			if err != nil {
				fmt.Println("Error in runner:", err)
				return
			}
			outputs = append(outputs, Output{ToolName: tool, Result: result})
		}(t)
	}
	wg.Wait()
	return outputs
}

func NewTestRunner(tools []string) *TestRunner {
	return &TestRunner{
		Tools: tools,
	}
}
