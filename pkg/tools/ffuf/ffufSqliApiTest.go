package ffuf

import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type FfufSqliApiTest struct {
	flags []string
	name  string
}

func (r FfufSqliApiTest) Execute(target string) (string, error) {
	output, err := runSqliApiTest(target, r)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func runSqliApiTest(target string, r FfufSqliApiTest) (string, error) {
	_, err := exec.LookPath("ffuf")
	if err != nil {
		utils.ToolFailed(r.name, target, err)
		return "", fmt.Errorf("ffuf not found")
	}
	utils.ToolStartLog(r.name, target)

	allFlags := append(r.flags, "-u", "http://"+target+"/rest/user/login", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"email":"FUZZ","password":"testtest"}`, "-ac", "-x", "http://localhost:9091")

	cmd := exec.Command("ffuf", allFlags...)
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

func NewSliApiTest(flags []string, name string) FfufSqliApiTest {
	return FfufSqliApiTest{flags: flags, name: name}
}
