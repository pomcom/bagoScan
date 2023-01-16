package nmap

import (
	"fmt"
	"os/exec"
	"strings"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type Nmap struct {
	flags []string
}

var tool = "nmap"

func (n Nmap) Execute(target string) (string, error) {

	output, err := runNmap(target, n)
	if err != nil {
		return "", err
	}

	fmt.Println(output)

	return output, nil
}

func (n Nmap) Name() string {
	return tool
}

func (n Nmap) SetFlags(flags []string) {
	n.flags = flags
	fmt.Println("set flags method:", flags)
}

func runNmap(target string, n Nmap) (string, error) {

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
	//cmd := exec.Command(tool, target)
	// cmd := exec.Command("nmap", append([]string{target}, n.flags...)...)
	println("flags nmap:", n.flags)
	// cmd := exec.Command("nmap", append([]string{target}, n.flags...)...)
	cmd := exec.Command("nmap", append(append([]string{}, n.flags...), target)...)

	fmt.Printf("Running command: %s %s\n", cmd.Path, strings.Join(cmd.Args[1:], " "))

	//Output() returns combined output of stdout and stderr
	//Seperation possible using StdoutPipe() and SterrPipe()
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
