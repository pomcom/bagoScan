package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

type Filehandler struct {
	outputDir string
}

func (handler *Filehandler) WriteToFile(filename string, output string) error {
	handler.outputDir = "output/raw"

	// create out/raw if it does not exist
	if err := handler.createDirectory(); err != nil {
		return err
	}

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

func (handler *Filehandler) createDirectory() error {
	if _, err := os.Stat(handler.outputDir); os.IsNotExist(err) {
		err = os.MkdirAll(handler.outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Error creating %s directory: %s", handler.outputDir, err)
		}
	}
	return nil
}
