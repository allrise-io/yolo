package cm

var TPL = `package cm

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func In(data string, dataList []string) bool {
	for _, item := range dataList {
		if data == item {
			return true
		}
	}
	return false
}

func Ini(num int, numList []int) bool {
	for _, n := range numList {
		if num == n {
			return true
		}
	}
	return false
}

func Call(param string) ([]byte, error) {
	c := exec.Command("/bin/bash", "-c", param)
	return c.CombinedOutput()
}

func Execute(param string) error {
	c := exec.Command("/bin/bash", "-c", param)
	return c.Run()
}

// fetch file path(go build)
func getExecFileAbPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	realPath, err := filepath.EvalSymlinks(filepath.Dir(execPath))
	if err != nil {
		return "", err
	}
	return realPath, nil
}

// fetch file path(go run)
func getDebugFileAbPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}

func GetProjectRootPath() (string, error) {
	dir, err := getExecFileAbPath()
	if err != nil {
		return "", err
	}
	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}
	if strings.Contains(dir, tmpDir) {
		curPath := getDebugFileAbPath()
		return filepath.Dir(filepath.Dir(filepath.Dir(curPath))), nil
	}
	return filepath.Dir(dir), nil
}

func TimeString(curTime time.Time) string {
	return curTime.Format("2006-01-02 15:04:05")
}
`
