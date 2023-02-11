package sqlmap

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type SQLMapApiTest struct {
	flags []string
	name  string
}

func (s SQLMapApiTest) Execute(target string) (string, error) {
	output, err := runSQLMapApiTest(target, s)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func NewSQLMapApiTest(options []string, name string) SQLMapApiTest {
	return SQLMapApiTest{flags: options, name: name}
}

func runSQLMapApiTest(target string, s SQLMapApiTest) (string, error) {

	_, err := exec.LookPath("sqlmap")
	if err != nil {
		utils.ToolFailed(s.name, target, err)
		return "", fmt.Errorf("sqlmap not found")
	}

	authToken, err := ioutil.ReadFile("auth_token.txt")
	if err != nil {
		utils.Logger.Warn("no auth token provided - starting sql withouth it")
	} else {
		s.flags = append(s.flags, fmt.Sprintf("--auth-token=%s", string(authToken)))
	}

	utils.ToolStartLog(s.name, target)

	//todo read in inject point from fuzzin phase, for protoytping currently like this:
	allFlags := append(s.flags, "-u", "http://"+target+"/rest/products/search?q=", "-D", "localhost", "--tables", "--level=5", "--batch")
	cmd := exec.Command("sqlmap", allFlags...)

	// i	allFlags := append(r.flags, "-u", "http://"+target+"/rest/user/login", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"email":"FUZZ","password":"testtest"}`, "-ac", "-x", "http://localhost:9091")
	//

	utils.ExecutedCommand(cmd)

	out, err := cmd.Output()
	if err != nil {
		utils.ToolFailed(s.name, target, err)
		return "", err
	}

	utils.ToolFinishedLog(s.name, target)

	return string(out), nil
}
