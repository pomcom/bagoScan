package ffuf

//needs to be adjusted for scalablity
import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type FfufBackupFiles struct {
	flags []string
	name  string
}

func (r FfufBackupFiles) Execute(target string) (string, error) {
	output, err := runBackupFilesFinder(target, r)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func runBackupFilesFinder(target string, r FfufBackupFiles) (string, error) {

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

func NewFfufBackupFiles(flags []string, name string) FfufBackupFiles {
	return FfufBackupFiles{flags: flags, name: name}
}
