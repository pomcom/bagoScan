package ffuf

import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type ResourceDiscovery struct {
	flags []string
	name  string
}

func (r ResourceDiscovery) Execute(target string) (string, error) {
	output, err := runDirBruteforce(target, r)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func runDirBruteforce(target string, r ResourceDiscovery) (string, error) {

	_, err := exec.LookPath("ffuf")

	if err != nil {
		utils.ToolFailed(r.name, target, err)
		return "", fmt.Errorf("ffuf not found")
	}

	utils.ToolStartLog(r.name, target)

	allFlags := append(r.flags, "-u", "http://"+target+"/FUZZ", "-fs", "1987")
	cmd := exec.Command("ffuf", allFlags...)

	// cmd := exec.Command("ffuf", append(r.flags, "-u", "http://"+target+"/FUZZ")...)
	// cmd := exec.Command("/home/pomcom/go/bin/ffuf", "-w", "common.txt", "--recursion-depth", "3", "-u", "http://"+target+"/FUZZ")
	// cmd.Env = append(cmd.Env, "PATH=$PATH:/usr/local/bin")

	println("running ffuf command:", cmd.String())

	utils.ExecutedCommand(cmd)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(string(out))
		println(err.Error, (string(out)))
		utils.ToolFailed(r.name, target, err)
		return "", err
	}

	fmt.Println(string(out))

	utils.ToolFinishedLog(r.name, target)
	return string(out), nil

}

func NewResourceDiscovery(flags []string, name string) ResourceDiscovery {
	return ResourceDiscovery{flags: flags, name: name}
}
