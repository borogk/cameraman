package main

import (
	"fmt"
	"os"

	"github.com/borogk/cameraman/internal/launch"
)

func main() {
	playerModulePath := launch.CameramanModulePath("CameramanPlayer.pk3")
	args := []string{"-file", playerModulePath}

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
}
