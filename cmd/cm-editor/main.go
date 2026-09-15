package main

import (
	"fmt"
	"os"

	"github.com/borogk/cameraman/launcher"
)

func main() {
	logFile, err := launcher.TempLogFile()
	if err != nil {
		fmt.Printf("error allocating temp log file: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		_ = logFile.Close()
	}()

	editorModulePath, err := launcher.CameramanModulePath("CameramanEditor.pk3")
	if err != nil {
		fmt.Printf("error finding PK3: %v\n", err)
		os.Exit(1)
	}

	args := []string{
		"-file", editorModulePath,
		"+freelook", "1",
		"+noclip2",
		"+notarget",
		"+logfile", logFile.Name(),
	}

	extraArgs, err := launcher.ParseExtraArgs()
	if err != nil {
		fmt.Printf("error parsing extra args: %v\n", err)
		os.Exit(1)
	}

	args = append(args, extraArgs...)
	err = launcher.LaunchCameraman(args)
	if err != nil {
		fmt.Printf("error running Cameraman: %s\n", err)
		os.Exit(1)
	}

	err = launcher.SaveExportedProfiles(logFile)
	if err != nil {
		fmt.Printf("error saving exported profiles: %s\n", err)
		os.Exit(1)
	}
}
