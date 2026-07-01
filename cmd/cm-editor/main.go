package main

import (
	"fmt"
	"os"

	"github.com/borogk/cameraman/internal/launch"
)

func main() {
	logFile := launch.TempLogFile()
	defer func() {
		_ = logFile.Close()
	}()

	editorModulePath := launch.CameramanModulePath("CameramanEditor.pk3")
	args := []string{
		"-file", editorModulePath,
		"+freelook", "1",
		"+noclip",
		"+notarget",
		"+logfile", logFile.Name(),
	}

	extraArgs, err := launch.ParseExtraArgs()
	if err != nil {
		fmt.Printf("error parsing extra args: %v\n", err)
		os.Exit(1)
	}

	args = append(args, extraArgs...)
	err = launch.RunCameraman(args)
	if err != nil {
		fmt.Printf("error running Cameraman: %s\n", err)
		os.Exit(1)
	}

	err = launch.SaveExportedProfiles(logFile)
	if err != nil {
		fmt.Printf("error saving exported profiles: %s\n", err)
		os.Exit(1)
	}
}
