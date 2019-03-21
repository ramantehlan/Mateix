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
	serverCmd.Bool("start", true, "Start the mateix server")
	serverCmd.Bool("stop", false, "stop the mateix server")

	// No sub-command
	if len(os.Args) < 2 {
		e.Error("Sub command required")
		os.Exit(1)
	}

	// Choose which sub-command
	switch os.Args[1] {
	case "update":
		updateCmd.Parse(os.Args[2:])

		// Update sub-command
		if updateCmd.Parsed() {
			if *updateFilePtr == "" {
				fmt.Println("usage of update:")
				updateCmd.PrintDefaults()
				os.Exit(1)
			}
			Update(*updateFilePtr)
		}

	case "init":
		Initialize()
	case "uninstall":
		Uninstall()
	case "server":
		serverCmd.Parse(os.Args[2:])

		if serverCmd.Parsed() && len(os.Args[0:]) <= 3 {
			arg := os.Args[2]
			if arg == "--start" || arg == "-start" {
				server(true)
			} else if arg == "--stop" || arg == "-stop" {
				server(false)
			} else {
				fmt.Println("usage of server:")
				serverCmd.PrintDefaults()
				os.Exit(1)
			}
		} else {
			fmt.Println("usage of server:")
			serverCmd.PrintDefaults()
			os.Exit(1)
		}

	default:
		e.Error("Unknown sub command")
		flag.PrintDefaults()
		os.Exit(1)
	}

}
