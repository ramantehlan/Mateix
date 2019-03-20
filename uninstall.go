package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
  "fmt"
)

func execute(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func getHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func uninstall() {
	var files = []string{
		"/usr/bin/mateix",
		"/usr/bin/mateixWatch",
		"/etc/systemd/system/multi-user.target.wants/mateix-watch.service",
		"/etc/systemd/system/mateix-watch.service",
		"/etc/.mateix",
		getHome() + "/.mateixConfig",
	}

  fmt.Println("Service stopped")
  execute(exec.Command("sudo", "mateixWatch", "stop"))

  for file , _ := range files {
		fi , err := os.Stat(files[file])
		if err == nil {
			switch mode := fi.Mode(); {
			case mode.IsDir():
				execute(exec.Command("sudo", "rm", "-r", files[file]))
        fmt.Println("Removed ", files[file])
			case mode.IsRegular():
				execute(exec.Command("sudo", "rm", files[file]))
        fmt.Println("Removed ", files[file])
			}
		}
	}


}
