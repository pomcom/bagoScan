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

func (n *Nmap) SetFlags(flags []string) {

	fmt.Println("set flags in set flags method:", flags)
	n.flags = flags
	fmt.Println("set flags in set flags method:", flags)
}

func runNmap(target string, n Nmap) (string, error) {

	// check if nmap is installed first
	_, err := exec.LookPath(tool)

	if err != nil {
		utils.ToolFailed(tool, target, err)
		return "", fmt.Errorf("nmap not found")
	}

	utils.ToolStartLog(tool, target)
	cmd := exec.Command("nmap", append(n.flags, target)...)
	println("flags nmap:", n.flags)
	fmt.Printf("Running command: %s %s\n", cmd.Path, strings.Join(cmd.Args[1:], " "))

	//Output() returns combined output of stdout and stderr
	//Seperation possible using StdoutPipe() and SterrPipe()
	out, err := cmd.Output()

	if err != nil {
		utils.ToolFailed(tool, target, err)
		return "", err
	}

	utils.ToolFinishedLog(tool, target)

	return string(out), nil
}
