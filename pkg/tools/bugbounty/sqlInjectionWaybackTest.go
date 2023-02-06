package bugbounty

import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type SqlInjectionWaybackTest struct {
	flags []string
	name  string
}

func (s SqlInjectionWaybackTest) Execute(target string) (string, error) {

	output, err := runSqlWaybackTest(target, s)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func NewSqlInjectionWayBackTest(flags []string, name string) SqlInjectionWaybackTest {
	return SqlInjectionWaybackTest{flags: flags, name: name}
}

func runSqlWaybackTest(target string, s SqlInjectionWaybackTest) (string, error) {

	utils.ToolStartLog("Sql Injection via Waybackurls", target)

	cmd := exec.Command("sh", "-c", "echo http://"+target+" | hakrawler -subs -u  | httpx -silent | anew | waybackurls | gf sqli >> sqli ; sqlmap -m sqli --batch --random-agent --level 5")

	utils.ExecutedCommand(cmd)
	println("running ffuf command:", cmd.String())
	out, err := cmd.Output()

	if err != nil {
		utils.ToolFailed(s.name, target, err)
		return "", err
	}

	utils.ToolFinishedLog("Sql Injection via Waybackurls", target)

	return string(out), nil
}
