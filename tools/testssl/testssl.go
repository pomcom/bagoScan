package testssl

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/pomcom/bagoScan/utils"
)

type Testssl struct{}

func (n *Testssl) Execute(flags string) (string, error) {

	output, err := scan(flags)
	if err != nil {
		return "", err
	}

	fmt.Println(output)

	fileHandler := &utils.Filehandler{}
	err = fileHandler.WriteToFile("testssl_output.txt", output)
	if err != nil {
		return "", fmt.Errorf("Error in testssl module writing output to file: %s", err)
	}
	return "", nil
}

func scan(target string) (string, error) {
	log.Println("Running testssl.sh on ", target)
	cmd := exec.Command("testssl.sh", target)
	out, err := cmd.Output()
	log.Println("testssl.sh finished.")
	if err != nil {
		return "", fmt.Errorf("Executing testssl failed with error: %s", err)
	}
	return string(out), nil
}
