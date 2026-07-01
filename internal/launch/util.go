package launch

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func CameramanModulePath(pk3 string) string {
	exec, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return path.Join(filepath.Dir(exec), pk3)
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

func TempLogFile() *os.File {
	logFile, err := os.CreateTemp(os.TempDir(), "*.log")
	if err != nil {
		panic(err)
	}
	return logFile
}
