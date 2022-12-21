package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	createFile()
	//TODO: populating output.txt with top metrics

	//TODO: exporting metrics to prometheus
}

func createFile() {
	usr, err := os.UserHomeDir()
	if err != nil {
		fmt.Print(err)
	}
	filePath := filepath.Join(usr, "Desktop")

	// createfile, err := os.OpenFile(filepath.Join(filePath, "output.txt"), os.O_RDWR|os.O_CREATE, 0644)

	createfile, err := os.Create(filepath.Join(filePath, "output.txt"))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(createfile)

}
