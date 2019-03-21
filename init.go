package main

import (
  "fmt"
  "os/exec"
  "github.com/ramantehlan/mateix/packages/command"
)

func Initialize() {
  fmt.Println("init command to list it in mateixWatch " + command.GetCurrentPath() )

  if ! command.FileExist(".mateix") {
    command.Execute(exec.Command("mkdir", ".mateix"))
  }


}
