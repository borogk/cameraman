package launcher

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
)

var exportFileNameRegexp = regexp.MustCompile("^export-([0-9]+)\\.cman$")

// ScanAndSaveExportedProfiles reads ZDoom's output and looks for camera profile exports.
func ScanAndSaveExportedProfiles(r io.Reader, outputDir string) (scanErr error) {
	var currentOutput *os.File
	defer func() {
		if currentOutput != nil {
			saveCurrentOutput(currentOutput)
			if scanErr == nil {
				scanErr = fmt.Errorf("last profile export was interrupted")
			}
		}
	}()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "--- BEGIN CAMERAMAN ---":
			if currentOutput != nil {
				return fmt.Errorf("unexpected profile begin before the previous export has ended")
			}
			output, err := os.Create(generateExportFileName(outputDir))
			if err != nil {
				return err
			}
			currentOutput = output
		case "--- END CAMERAMAN ---":
			if currentOutput != nil {
				saveCurrentOutput(currentOutput)
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

// generateExportFileName allocates a new available export file name (export-0001.cman, export-0002.cman etc.).
func generateExportFileName(outputDir string) string {
	cmanFiles, _ := filepath.Glob(path.Join(outputDir, "*.cman"))

	maxNum := 0
	for _, file := range cmanFiles {
		file := filepath.Base(file)
		submatch := exportFileNameRegexp.FindStringSubmatch(file)
		if len(submatch) == 2 {
			num, _ := strconv.ParseInt(submatch[1], 10, 32)
			maxNum = max(maxNum, int(num))
		}
	}

	return path.Join(outputDir, fmt.Sprintf("export-%04d.cman", maxNum+1))
}

func saveCurrentOutput(currentOutput *os.File) {
	fmt.Printf("saved %s\n", currentOutput.Name())
	_ = currentOutput.Close()
}
