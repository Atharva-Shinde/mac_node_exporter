package extractor

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Extractor() runs ./script.sh file and converts the result into float64
func Extractor() float64 {

	cmd := exec.Command("/bin/bash", "./script.sh")
	rawoutput, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
	}
	stroutput := string(rawoutput)
	finaloutput := strings.TrimSpace(stroutput)

	result, err := strconv.ParseFloat(finaloutput, 64)
	if err != nil {
		panic(err)
	}
	// fmt.Print(result)
	return result
}
