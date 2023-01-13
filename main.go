package main

import (
	"fmt"

	"github.com/pomcom/bagoScan/commands"
)

func main() {

	err := commands.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}

}
