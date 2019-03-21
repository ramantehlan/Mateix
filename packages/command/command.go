package command

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/ramantehlan/mateix/packages/e"
)

// Execute is to execute a command in terminal
func Execute(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	e.Check(err)
}

// GetHome is to get the home directory
func GetHome() string {
	usr, err := user.Current()
	e.Check(err)
	return usr.HomeDir
}

// GetUser is to get the current user
func GetUser() string {
	usr, err := user.Current()
	e.Check(err)
	return usr.Username
}

// GetCurrentPath is to get the current path
func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	e.Check(err)
	return dir
}

// GetStat is to get the current status of a file
func GetStat(path string) os.FileInfo {
	fi, err := os.Stat(path)
	e.Check(err)
	return fi
}

// FileExist is to check if a file exist
func FileExist(path string) bool {
	fileExist := false
	_, err := os.Stat(path)
	if err == nil {
		fileExist = true
	}
	return fileExist
}
