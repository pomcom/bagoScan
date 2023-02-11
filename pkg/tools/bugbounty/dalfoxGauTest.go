package bugbounty

import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type DalfoxGauTest struct {
	flags []string
	name  string
}

func (d DalfoxGauTest) Execute(target string) (string, error) {

	output, err := runDalfoxGauTest(target, d)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func NewDalfoxGauTest(flags []string, name string) DalfoxGauTest {
	return DalfoxGauTest{flags: flags, name: name}
}

func runDalfoxGauTest(target string, s DalfoxGauTest) (string, error) {

	utils.ToolStartLog("XSS Test using  dalfox and gau", target)

	cmd := exec.Command("sh", "-c", "assetfinder "+target+" | gau | dalfox pipe")
	// cmd := exec.Command("assetfinder", target, "|", "gau", "|", "dalfox", "pipe")
	//assetfinder testphp.vulnweb.com | gau | dalfox pipe
	utils.ExecutedCommand(cmd)
	println("running ffuf command:", cmd.String())
	out, err := cmd.Output()

	if err != nil {
		utils.ToolFailed(s.name, target, err)
		return "", err
	}

	utils.ToolFinishedLog("XSS Test using dalfox", target)

	return string(out), nil
}
