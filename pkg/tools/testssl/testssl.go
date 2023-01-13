package testssl

import (
	"fmt"
	"log"
	"os/exec"
)

type Testssl struct{}

func (t Testssl) Execute(target string) (string, error) {

	output, err := scan(target)
	if err != nil {
		return "", err
	}

	fmt.Println(output)

	if err != nil {
		return output, fmt.Errorf("Error in testssl module writing output to file: %s", err)
	}
	return output, nil
}

func (t Testssl) Name() string {
	return "testssl"
}

func scan(target string) (string, error) {
	log.Println("Running testssl.sh on", target)
	cmd := exec.Command("testssl.sh", target)
	out, err := cmd.Output()
	log.Println("testssl.sh finished.")
	if err != nil {
		return "", fmt.Errorf("Executing testssl failed with error: %s", err)
	}
	return string(out), nil
}
