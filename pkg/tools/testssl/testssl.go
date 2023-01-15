package testssl

import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type Testssl struct{}

var tool = "testssl.sh"

func (t Testssl) Execute(target string) (string, error) {

	output, err := scan(target)
	if err != nil {
		return "", err
	}

	fmt.Println(output)

	return output, nil
}

func (t Testssl) Name() string {
	return tool
}

func scan(target string) (string, error) {

	utils.ToolStartLog(tool, target)

	cmd := exec.Command(tool, target)
	out, err := cmd.Output()

	utils.ToolFinishedLog(tool, target)

	if err != nil {
		utils.ToolFailed(tool, target, err)
		return "", err
	}
	return string(out), nil
}
