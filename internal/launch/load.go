package launch

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var cvarRegexp = regexp.MustCompile("^(\\w+) = (.*)$")
var allowedCvarNames = make(map[string]bool)
var allowedCvarNamesRaw = []string{
	"path_mode",
	"speed_mode",
	"angle_mode",
	"delay",
	"speed",
	"overshoot",
	"warp_player",
	"hide_player",
	"ga_buffer_len",
	"x0",
	"x1",
	"x2",
	"y0",
	"y1",
	"y2",
	"z0",
	"z1",
	"z2",
	"a0",
	"a1",
	"p0",
	"p1",
	"ra0",
	"ra1",
	"cx0",
	"cx1",
	"cy0",
	"cy1",
	"r0",
	"r1",
}

func init() {
	for _, name := range allowedCvarNamesRaw {
		allowedCvarNames[name] = true
	}
}

func ParseExtraArgs() ([]string, error) {
	if len(os.Args) >= 2 && os.Args[1] == "load" {
		if len(os.Args) < 3 {
			return nil, fmt.Errorf("missing filename after 'load' command")
		}

		cmanProfilePath := os.Args[2]
		cmanArgs, err := ParseCameraProfile(cmanProfilePath)
		if err != nil {
			return nil, err
		}

		userArgs := os.Args[3:]
		return append(userArgs, cmanArgs...), nil
	}

	userArgs := os.Args[1:]
	return userArgs, nil
}

func ParseCameraProfile(filePath string) ([]string, error) {
	fmt.Println("loading CVARs...")

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	args := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		submatch := cvarRegexp.FindStringSubmatch(scanner.Text())
		if len(submatch) == 3 {
			name := submatch[1]
			value := submatch[2]
			if allowedCvarNames[name] {
				args = append(args, "+cman_"+name, value)
			} else {
				fmt.Printf("skipping non-Cameraman CVAR %s\n", name)
			}
		}
	}

	return args, nil
}
