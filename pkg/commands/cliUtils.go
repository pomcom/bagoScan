// Package commands does something
package commands

// cli related utils, used in multiple commands
import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Filter out any empty strings in target slice before passing it
// otherwise empty strings get used as a target
func filterEmptyStrings(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func checkFlags() {
	if len(target) == 0 && targetFile == "" {
		fmt.Println("Error: either --target or --target-file must be provided")
		return
	}
}

func getTargets(targetFile string, target []string) []string {
	if targetFile != "" {
		targetsBytes, err := ioutil.ReadFile(targetFile)
		if err != nil {
			fmt.Printf("Error reading target file: %v", err)
			return nil
		}
		targets := strings.Split(string(targetsBytes), "\n")
		targets = filterEmptyStrings(targets, func(i string) bool {
			return i != ""
		})
		return targets
	}
	return target
}
