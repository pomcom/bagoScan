package main

import (
	"fmt"
	"log"

	"github.com/pomcom/bagoScan/tools/testssl"
)

func main() {

	log.Println("Starting ...")

	t := &testssl.Testssl{}
	_, err := t.Execute("pomcom.digital")
	if err != nil {
		fmt.Println(err)
		return
	}

}
