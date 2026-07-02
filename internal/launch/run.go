package launch

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunCameraman(args []string) error {
	zdoomBin := ZdoomBin()
	fmt.Printf("ZDoom binary: %s\n", zdoomBin)
	fmt.Printf("ZDoom args: %s\n", strings.Join(args, " "))
	fmt.Println("Cameraman is launching...")

	err := exec.Command(zdoomBin, args...).Run()
	if err != nil {
		return err
	}

	return nil
}
