package utils

import (
	"fmt"
	"sync"

	"github.com/pomcom/bagoScan/pkg/tools"
)

type Output struct {
	ToolName string
	Result   string
}

type TestRunner struct {
	ToolMap map[string]tools.Tool
}

func (r TestRunner) Run(target string) []Output {
	var wg sync.WaitGroup
	wg.Add(len(r.ToolMap))
	outputs := make([]Output, 0)

	for toolName := range r.ToolMap {
		go func(toolName string) {
			defer wg.Done()
			tool := r.ToolMap[toolName]
			result, err := tool.Execute(target)
			if err != nil {
				fmt.Println("Error in runner:", err)
				return
			}
			outputs = append(outputs, Output{ToolName: toolName, Result: result})
		}(toolName)
	}
	wg.Wait()
	return outputs
}

func NewTestRunner(toolMap map[string]tools.Tool) TestRunner {
	return TestRunner{
		ToolMap: toolMap,
	}
}
