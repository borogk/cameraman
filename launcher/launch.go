package launcher

import (
	"fmt"
	"os/exec"
)

// LaunchCameraman starts ZDoom executable and optionally scans stdout for camera profile exports.
func LaunchCameraman(args []string, scanForExports bool) error {
	zdoomBin := ZdoomBin()
	fmt.Printf("ZDoom binary: \n")
	fmt.Printf("  %s\n", zdoomBin)
	fmt.Printf("ZDoom args: \n")
	for _, arg := range args {
		fmt.Printf("  %s\n", arg)
	}

	fmt.Println("-------------------------")
	fmt.Println("Cameraman is launching...")
	defer func() {
		fmt.Println("Cameraman session completed")
	}()

	cmd := exec.Command(zdoomBin, args...)
	if scanForExports {
		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}

		err = cmd.Start()
		if err != nil {
			return err
		}

		err = ScanAndSaveExportedProfiles(stdoutPipe)
		if err != nil {
			fmt.Printf("error while scanning for exported profiles: %v\n", err)
		}

		return cmd.Wait()
	}

	return cmd.Run()
}
