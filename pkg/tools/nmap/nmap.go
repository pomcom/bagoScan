package nmap

import (
	"fmt"
	"os/exec"

	"github.com/pomcom/bagoScan/pkg/utils"
)

type Nmap struct{}

var tool = "nmap"

func (n Nmap) Execute(target string) (string, error) {

	output, err := scan(target)
	if err != nil {
		return "", err
	}

	fmt.Println(output)

	if err != nil {
		return output, fmt.Errorf("Error in nmap module writing output to file: %s", err)
	}
	return output, nil
}

func (n Nmap) Name() string {
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
