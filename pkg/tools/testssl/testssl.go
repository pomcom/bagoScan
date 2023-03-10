package testssl

import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type Testssl struct {
	flags []string
	name  string
}

func (t Testssl) Execute(target string) (string, error) {

	output, err := runTessl(target, t)
	if err != nil {
		return "", err
	}

	fmt.Println(output)
	return output, nil
}

func NewTestssl(flags []string, name string) Testssl {
	return Testssl{flags: flags, name: name}
}

func runTessl(target string, t Testssl) (string, error) {

	utils.ToolStartLog(t.name, target)

	cmd := exec.Command("testssl.sh", append(t.flags, target)...)
	utils.ExecutedCommand(cmd)

	out, err := cmd.Output()

	utils.ToolFinishedLog(t.name, target)

	if err != nil {
		utils.ToolFailed(t.name, target, err)
		return "", err
	}
	return string(out), nil
}
