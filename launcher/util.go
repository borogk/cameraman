package launcher

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func CameramanModulePath(pk3 string) (string, error) {
	exec, err := os.Executable()
	if err != nil {
		return "", err
	}

	return path.Join(filepath.Dir(exec), pk3), nil
}

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

func TempLogFile() (*os.File, error) {
	logFile, err := os.CreateTemp(os.TempDir(), "*.log")
	if err != nil {
		return nil, err
	}

	return logFile, nil
}
