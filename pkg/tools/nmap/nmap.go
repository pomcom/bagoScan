package nmap

import (
	"fmt"
	"os/exec"
	"strings"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
	"go.uber.org/zap"
)

type Nmap struct {
	flags []string
	name  string
}

func (n Nmap) Execute(target string) (string, error) {

	output, err := runNmap(target, n)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func NewNmap(flags []string, name string) Nmap {
	return Nmap{flags: flags, name: name}
}

func runNmap(target string, n Nmap) (string, error) {

	// check if nmap is installed first
	_, err := exec.LookPath(n.name)

	if err != nil {
		utils.ToolFailed(n.name, target, err)
		return "", fmt.Errorf("nmap not found")
	}

	utils.ToolStartLog(n.name, target)
	cmd := exec.Command("nmap", append(n.flags, target)...)
	//	fmt.Printf("Running command: %s %s\n", cmd.Path, strings.Join(cmd.Args[1:], " "))
	utils.Logger.Info("Full command", zap.String(" ", cmd.Path), zap.String(" ", strings.Join(cmd.Args[1:], "")))

	//Output() returns combined output of stdout and stderr
	//Seperation possible using StdoutPipe() and SterrPipe()
	out, err := cmd.Output()

	if err != nil {
		utils.ToolFailed(n.name, target, err)
		return "", err
	}

	utils.ToolFinishedLog(n.name, target)

	return string(out), nil
}
