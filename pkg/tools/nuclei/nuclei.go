package nuclei

import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type Nuclei struct {
	flags []string
	name  string
}

func (n Nuclei) Execute(target string) (string, error) {
	output, err := runNuclei(target, n)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func NewNuclei(flags []string, name string) Nuclei {
	return Nuclei{flags: flags, name: name}
}

func runNuclei(target string, n Nuclei) (string, error) {
	_, err := exec.LookPath(n.name)
	if err != nil {
		utils.ToolFailed(n.name, target, err)
		return "", fmt.Errorf("nuclei not found")
	}

	utils.ToolStartLog(n.name, target)
	cmd := exec.Command("nuclei", append(n.flags, target)...)
	utils.ExecutedCommand(cmd)
	out, err := cmd.Output()
	if err != nil {
		utils.ToolFailed(n.name, target, err)
		return "", err
	}
	utils.ToolFinishedLog(n.name, target)
	return string(out), nil
}
