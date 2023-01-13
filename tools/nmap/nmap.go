package nmap

import (
	"fmt"
	"log"
	"os/exec"
)

type Nmap struct{}

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
	return "nmap"
}

func scan(target string) (string, error) {
	log.Println("Running nmap on", target)
	cmd := exec.Command("nmap", target)
	out, err := cmd.Output()
	log.Println("nmap finished.")
	if err != nil {
		return "", fmt.Errorf("Executing nmap failed with error: %s", err)
	}
	return string(out), nil
}
