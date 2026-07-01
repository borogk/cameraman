package launch

import (
	"fmt"
	"os/exec"
)

func RunCameraman(args []string) error {
	zdoomBin := ZdoomBin()
	fmt.Printf("ZDoom binary: %s\n", zdoomBin)
	fmt.Printf("ZDoom args: %v\n", args)
	fmt.Println("Cameraman is launching...")

	err := exec.Command(zdoomBin, args...).Run()
	if err != nil {
		return err
	}

	return nil
}
