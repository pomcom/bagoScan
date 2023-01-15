package nmap

import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type Nmap struct{}

var tool = "nmap"

func (n Nmap) Execute(target string) (string, error) {

	output, err := runNmap(target)
	if err != nil {
		return "", err
	}

	fmt.Println(output)

	return output, nil
}

func (n Nmap) Name() string {
	return tool
}

func runNmap(target string) (string, error) {

	// check if nmap is installed first
	_, err := exec.LookPath(tool)

	if err != nil {
		utils.ToolFailed(tool, target, err)
		return "", fmt.Errorf("nmap not found")
	}

	// check if target is reachable - check if ping is in path?
	pingCmd := exec.Command("ping", "-c 1", "-W 1", target)
	if err := pingCmd.Run(); err != nil {
		return "", fmt.Errorf("target %s is not reachable", target)
	}

	utils.ToolStartLog(tool, target)
	cmd := exec.Command(tool, target)
	out, err := cmd.Output()
	if err != nil {
		utils.ToolFailed(tool, target, err)
		return "", err
	}

	utils.ToolFinishedLog(tool, target)

	if err != nil {
		utils.ToolFailed(tool, target, err)
		return "", err
	}
	return string(out), nil
}
