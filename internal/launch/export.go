package launch

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var exportFileNameRegexp = regexp.MustCompile("^export-([0-9]+)\\.cman$")

func SaveExportedProfiles(r io.Reader) error {
	var currentOutput *os.File
	defer func() {
		if currentOutput != nil {
			_ = currentOutput.Close()
		}
	}()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "--- BEGIN CAMERAMAN ---":
			output, err := os.Create(generateExportFileName())
			if err != nil {
				return err
			}
			currentOutput = output
		case "--- END CAMERAMAN ---":
			if currentOutput != nil {
				fmt.Printf("saving %s\n", currentOutput.Name())
				_ = currentOutput.Close()
				currentOutput = nil
			}
		default:
			if currentOutput != nil {
				_, err := currentOutput.WriteString(line + "\n")
				if err != nil {
					return err
				}
			}
		}
	}

	return scanner.Err()
}

func generateExportFileName() string {
	cmanFiles, _ := filepath.Glob("*.cman")

	maxNum := 0
	for _, file := range cmanFiles {
		submatch := exportFileNameRegexp.FindStringSubmatch(file)
		if len(submatch) == 2 {
			num, _ := strconv.ParseInt(submatch[1], 10, 32)
			maxNum = max(maxNum, int(num))
		}
	}

	return fmt.Sprintf("export-%04d.cman", maxNum+1)
}
