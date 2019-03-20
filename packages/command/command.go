package command

import (
  "os"
	"os/user"
  "os/exec"
  "path/filepath"
  "github.com/ramantehlan/mateix/packages/error"
  )

func Execute(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	error.Check(err)
}

func GetHome() string {
	usr, err := user.Current()
	error.Check(err)
	return usr.HomeDir
}

func GetCurrentPath() string {
  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  error.Check(err)
  return dir
}

func GetStat(path string) os.FileInfo {
  fi , err := os.Stat(path)
  error.Check(err)
  return fi
}

func FileExist(path string) bool{
  fileExist := false
  _ , err := os.Stat(path)
  if err == nil{
    fileExist = true
  }

  return fileExist
}
