package main

import (
	"flag"
	"fmt"
	"os"
  "github.com/ramantehlan/mateix/packages/error"
)

func main() {

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateFilePtr := updateCmd.String("file", "", "File path which you wish to sync")

	// No sub-command
	if len(os.Args) < 2 {
		error.Error("Sub command required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "update":
		updateCmd.Parse(os.Args[2:])
  case "init":
    fmt.Println("init command to list it in mateixWatch")
  case "uninstall":
    fmt.Println("To uninstall the program")
	default:
		error.Error("Unknown sub command")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if updateCmd.Parsed() {
		if *updateFilePtr == "" {
      fmt.Println("usage of update:")
			updateCmd.PrintDefaults()
			os.Exit(1)
		}
		// Print
		fmt.Println("You are here")
	}

}
