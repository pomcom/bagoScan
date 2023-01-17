package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Filehandler struct {
	outputDir string
}

func NewFilehandler(outputDir string) Filehandler {
	// set default path, if no path is provided
	if outputDir == "" {
		outputDir = "output/raw"
	}
	return Filehandler{outputDir: outputDir}
}

func (handler Filehandler) WriteToFile(filename string, output string) error {
	// create out/raw if it does not exist
	if err := handler.createDirectory(); err != nil {
		return err
	}

	//add timestamp
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	filename = fmt.Sprintf("%s-%s", timestamp, filename)

	// create file in output/raw
	filePath := filepath.Join(handler.outputDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error creating file: %s", err)
	}
	defer file.Close()

	_, err = file.WriteString(output)

	if err != nil {
		return fmt.Errorf("Error writing to file: %s", err)

	}
	return nil
}

/*
Could be better to adjust the outputDir creation,
since it already does get created in main.go when in initializing the
utils logger
*/
func (handler *Filehandler) createDirectory() error {
	if _, err := os.Stat(handler.outputDir); os.IsNotExist(err) {
		err = os.MkdirAll(handler.outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Error creating %s directory: %s", handler.outputDir, err)
		}
	}
	return nil
}