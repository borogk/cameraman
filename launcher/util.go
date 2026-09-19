package launcher

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// CameramanModulePath returns the absolute path of a .pk3 module that is located next to Cameraman executable.
func CameramanModulePath(pk3 string) (string, error) {
	exec, err := os.Executable()
	if err != nil {
		return "", err
	}

	return path.Join(filepath.Dir(exec), pk3), nil
}

// ZdoomBin returns the path to configured ZDoom executable.
func ZdoomBin() string {
	env := os.Getenv("ZDOOM_BIN")
	if env != "" {
		return env
	}

	switch runtime.GOOS {
	case "windows":
		return "uzdoom.exe"
	default:
		return "uzdoom"
	}
}
