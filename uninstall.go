package main

import (
	"os/exec"
  "fmt"
  "github.com/ramantehlan/mateix/packages/command"
)

func uninstall() {
	var files = []string{
		"/usr/bin/mateix",
		"/usr/bin/mateixWatch",
		"/etc/systemd/system/multi-user.target.wants/mateix-watch.service",
		"/etc/systemd/system/mateix-watch.service",
		"/etc/.mateix",
		command.GetHome() + "/.mateixConfig",
	}

  fmt.Println("Service stopped")
  command.Execute(exec.Command("sudo", "mateixWatch", "stop"))

  for file , _ := range files {
		if command.FileExist(files[file]) {
			fi := command.GetStat(files[file])
			switch mode := fi.Mode(); {
			case mode.IsDir():
				command.Execute(exec.Command("sudo", "rm", "-r", files[file]))
        fmt.Println("Removed ", files[file])
			case mode.IsRegular():
				command.Execute(exec.Command("sudo", "rm", files[file]))
        fmt.Println("Removed ", files[file])
			}
		}
	}

}
