package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	createFile()
}

func createFile() {
	usr, err := os.UserHomeDir()
	if err != nil {
		fmt.Print(err)
	}
	filePath := filepath.Join(usr, "Desktop")
	os.MkdirAll(filePath, 0755)
	createfile, err := os.OpenFile(filepath.Join(filePath, "output.txt"), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(createfile)
}
