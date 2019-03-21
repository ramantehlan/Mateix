package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ramantehlan/mateix/packages/e"
)

func main() {

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateFilePtr := updateCmd.String("file", "", "File path which you wish to sync")
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverStartPtr := serverCmd.Bool("start", true, "Start the mateix server")
	serverStopPtr := serverCmd.Bool("stop", true, "Start the mateix server")

	// No sub-command
	if len(os.Args) < 2 {
		e.Error("Sub command required")
		os.Exit(1)
	}

	// Choose which sub-command
	switch os.Args[1] {
	case "update":
		updateCmd.Parse(os.Args[2:])
	case "init":
		Initialize()
	case "uninstall":
		Uninstall()
	case "server":
		serverCmd.Parse(os.Args[2:])
	default:
		e.Error("Unknown sub command")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if serverCmd.Parsed() {
		if *serverStartPtr {
			server(true)
		} else if *serverStopPtr {
			server(false)
		}
	}

	// Update sub-command
	if updateCmd.Parsed() {
		if *updateFilePtr == "" {
			fmt.Println("usage of update:")
			updateCmd.PrintDefaults()
			os.Exit(1)
		}
		Update(*updateFilePtr)
	}

}
