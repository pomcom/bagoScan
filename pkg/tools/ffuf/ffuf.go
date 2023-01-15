package ffuf

import (
	"os/exec"
)

type Ffuf struct {
	options []string
}

var tool = "ffuf"

func (f Ffuf) Name() string {
	return tool
}
func (f Ffuf) AddOption(opt string) {
	f.options = append(f.options, opt)
}

func (f Ffuf) Execute(target string) (string, error) {
	cmd := exec.Command("ffuf", append([]string{target}, f.options...)...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
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
