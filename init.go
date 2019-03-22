package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/ramantehlan/mateix/packages/command"
	"github.com/ramantehlan/mateix/packages/e"
)

// Config is the structure of any config file
type Config struct {
	TargetIP  string `json:"targetIp"`
	TargetDir string `json:"entity"`
}

/*
"targetUser": "mspp",
"targetIp": "192.168.43.189",
"targetDir": "/home/mspp/mateixTest"
*/

// CreateJSON is to create a json file
func CreateJSON(config Config, jsonFile string) {
	file, err := json.MarshalIndent(config, "", " ")
	e.Check(err)
	err = ioutil.WriteFile(jsonFile, file, 0644)
	e.Check(err)
}

// Initialize the mateix watch folder
func Initialize() {
	u, _ := user.Current()

	if u.Uid == "0" {

		if !command.FileExist(".mateix") {
			command.Execute(exec.Command("sudo", "mateixWatch", "stop"))
			fmt.Println("MateixWatch Service Stopped")
			conf := Config{}
			fmt.Println("This will initialize mateix watch in this folder:")
			fmt.Print("Target IP: ")
			reader := bufio.NewReader(os.Stdin)
			conf.TargetIP, _ = reader.ReadString('\n')
			conf.TargetIP = strings.ReplaceAll(conf.TargetIP, "\n", "")
			fmt.Print("Target Directory: ")
			conf.TargetDir, _ = reader.ReadString('\n')
			conf.TargetDir = strings.ReplaceAll(conf.TargetDir, "\n", "")

			dat, err := ioutil.ReadFile("/etc/.mateix/syncList")
			e.Check(err)
			text := string(dat) + command.GetCurrentPath() + "\n"
			syncList := []byte(text)
			ioutil.WriteFile("/etc/.mateix/syncList", syncList, 0644)
			command.Execute(exec.Command("mkdir", command.GetCurrentPath()+"/.mateix"))
			command.Execute(exec.Command("touch", command.GetCurrentPath()+"/.mateix/config.json"))
			CreateJSON(conf, command.GetCurrentPath()+"/.mateix/config.json")

			fmt.Printf("Added '%s' in '/etc/.mateix/syncList'\n", command.GetCurrentPath())
			fmt.Println("Created ", command.GetCurrentPath()+"/.mateix")
			fmt.Println("Created ", command.GetCurrentPath()+"/.mateix/config.json")
			fmt.Println("MateixWatch Service Started Again")

			command.Execute(exec.Command("sudo", "mateixWatch", "start"))

		} else {
			fmt.Printf("%s is already a mateix watched file\n", command.GetCurrentPath())
		}

	} else {
		fmt.Println("You need to root to execute this command")
	}

}
