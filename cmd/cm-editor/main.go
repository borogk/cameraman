package main

import (
	"fmt"
	"os"

	"github.com/borogk/cameraman/launcher"
)

func main() {
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
		"+god",
		"+r_drawplayersprites", "0",
		"-stdout",
	}

	extraArgs, err := launcher.ParseLaunchArgs(os.Args, "Cman_WarpToPath")
	if err != nil {
		fmt.Printf("error parsing extra args: %v\n", err)
		os.Exit(1)
	}

	args = append(args, extraArgs...)
	err = launcher.LaunchCameraman(args, true)
	if err != nil {
		fmt.Printf("error running Cameraman: %v\n", err)
		os.Exit(1)
	}
}
