package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	t.Run("Testing for created file", func(t *testing.T) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			t.Error(err)
		}
		filePath := homeDir + "Desktop/output.txt"
		// got := path.Base(hd)
		check, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			fmt.Print("file doesn't exist")
		} else if err != nil {
			t.Error(err)
		} else {
			fmt.Print("file exists")
		}
		fmt.Print(check)
	})
}
