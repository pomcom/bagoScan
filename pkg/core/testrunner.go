package core

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

// tried to limit the number of go routines for performance handling
// https://granulate.io/blog/deep-dive-into-golang-performance/
// currently waiting till function is finished, before it returns
// could write output to filesystem in routins (bad design)
// or send output via channel to handler (could be timeconsuming to implement, but possible)
func (runner TestRunner) Run(targets []string) []Output {
	var wg sync.WaitGroup
	var outputs []Output
	//limit routines
	semaphore := make(chan struct{}, 150)

	for _, target := range targets {
		for toolName := range runner.ToolMap {
			semaphore <- struct{}{}
			wg.Add(1)
			go func(toolName, target string) {
				defer wg.Done()

				defer func() { <-semaphore }()
				tool := runner.ToolMap[toolName]
				result, err := tool.Execute(target)
				if err != nil {
					fmt.Println("Error in runner:", err)
					return
				}
				outputs = append(outputs, Output{ToolName: toolName, Result: result, Target: target})
			}(toolName, target)
		}
	}
	wg.Wait()
	return outputs
}

func NewTestRunner(toolMap map[string]tools.Tool) TestRunner {
	return TestRunner{
		ToolMap: toolMap,
	}
}
