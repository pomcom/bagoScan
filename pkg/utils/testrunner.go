package utils

import (
	"fmt"
	"sync"

	"github.com/pomcom/bagoScan/pkg/tools"
)

type Output struct {
	ToolName string
	Result   string
	Target   string
}

type TestRunner struct {
	ToolMap map[string]tools.Tool
}

func (runner TestRunner) Run(target string) []Output {
	var wg sync.WaitGroup
	wg.Add(len(runner.ToolMap))
	outputs := make([]Output, 0)

	for toolName := range runner.ToolMap {
		go func(toolName string) {
			defer wg.Done()
			tool := runner.ToolMap[toolName]
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
