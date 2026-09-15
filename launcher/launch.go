package launcher

import (
	"fmt"
	"os/exec"
)

func LaunchCameraman(args []string) error {
	zdoomBin := ZdoomBin()
	fmt.Printf("ZDoom binary: \n")
	fmt.Printf("  %s\n", zdoomBin)
	fmt.Printf("ZDoom args: \n")
	for _, arg := range args {
		fmt.Printf("  %s\n", arg)
	}
	fmt.Printf("\n")

	fmt.Println("Cameraman is launching...")
	err := exec.Command(zdoomBin, args...).Run()
	if err != nil {
		return err
	}

	return nil
}
