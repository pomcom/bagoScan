package ffuf

import (
	"fmt"
	"os/exec"

	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
)

type Ressource_discovery struct {
	options []string
}

var tool = "ffuf"

func (f Ressource_discovery) Name() string {
	return tool
}
func (f Ressource_discovery) AddOption(opt string) {
	f.options = append(f.options, opt)
}

func (f Ressource_discovery) Execute(target string) (string, error) {

	options := []string{"-w", "resources/common.txt", "-u", "http://" + target + "/FUZZ"}
	utils.ToolStartLog(tool, target)
	cmd := exec.Command("ffuf", options...)
	out, err := cmd.Output()
	if err != nil {
		utils.ToolFailed(tool, target, err)
		return "", err
	}
	fmt.Println(string(out))

	utils.ToolFinishedLog(tool, target)
	return string(out), nil
}

/*
Verschiedene Fuff automation = verschiedene "Tools" unabhaenig voneinander?
Wordlist übergeben usw  - erst customs flags implementieren?
Wordlists,
einfach in Tool einbinden unter Ressources? -
Public helpers mit set usw ? - Nein
Lib support ? Nein :(
struct Ffuf resource discvoery
*/
