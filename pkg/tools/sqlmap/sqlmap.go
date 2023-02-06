package sqlmap

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type SQLMap struct {
	options []string
	name    string
}

func (s SQLMap) Execute(target string) (string, error) {
	output, err := runSQLMap(target, s)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func NewSQLMap(options []string, name string) SQLMap {
	return SQLMap{options: options, name: name}
}

func runSQLMap(target string, s SQLMap) (string, error) {
	_, err := exec.LookPath(s.name)
	if err != nil {
		utils.ToolFailed(s.name, target, err)
		return "", fmt.Errorf("sqlmap not found")
	}

	authToken, err := ioutil.ReadFile("auth_token.txt")
	if err != nil {
		utils.Logger.Warn("no auth token provided - starting sql withouth it")
	} else {
		s.options = append(s.options, fmt.Sprintf("--auth-token=%s", string(authToken)))
	}

	utils.ToolStartLog(s.name, target)
	cmd := exec.Command("sqlmap", append(s.options, target)...)
	utils.ExecutedCommand(cmd)

	out, err := cmd.Output()
	if err != nil {
		utils.ToolFailed(s.name, target, err)
		return "", err
	}

	utils.ToolFinishedLog(s.name, target)

	return string(out), nil
}
