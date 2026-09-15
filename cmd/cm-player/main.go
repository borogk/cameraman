package main

import (
	"fmt"
	"os"

	"github.com/borogk/cameraman/launcher"
)

func main() {
	playerModulePath, err := launcher.CameramanModulePath("CameramanPlayer.pk3")
	if err != nil {
		fmt.Printf("error finding PK3: %v\n", err)
		os.Exit(1)
	}

	args := []string{"-file", playerModulePath}

	extraArgs, err := launcher.ParseLaunchArgs("Cman_PlayInPlayer")
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
}
