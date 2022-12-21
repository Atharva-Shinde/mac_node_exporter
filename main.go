package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	createFile()

	//TODO: exporting metrics to prometheus
}

func createFile() {
	usr, err := os.UserHomeDir()
	if err != nil {
		fmt.Print(err)
	}
	filePath := filepath.Join(usr, "Desktop")

	// createfile, err := os.OpenFile(filepath.Join(filePath, "output.txt"), os.O_RDWR|os.O_CREATE, 0644)
	absPath := filepath.Join(filePath, "output.txt")
	createfile, err := os.Create(absPath)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(createfile)

	populater(usr, filePath, absPath)
}

func populater(usr, filePath, absPath string) {
	getstats := exec.Command("top", "-l1", "-n10", "-stats", "pid,cpu")
	output, err := getstats.CombinedOutput()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(output))

	//sed can also be used to delete the summary

	populate := os.WriteFile(absPath, output, 0755)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(populate)

}
