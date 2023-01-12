package testssl

import (
	"fmt"
	"os/exec"
)

type Testssl struct{}

func (n *Testssl) Execute(flags string) (string, error) {

	output, err := scan(flags)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func scan(target string) (string, error) {
	cmd := exec.Command("testssl.sh", target)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("Executing testssl failed with error: %s", err)
	}
	return string(out), nil
}
