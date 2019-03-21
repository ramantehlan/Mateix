package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/ramantehlan/mateix/packages/e"
)

func connectToServer(addr string, port int) net.Conn {
	dest := addr + ":" + strconv.Itoa(port)
	fmt.Printf("Connecting to %s...\n", dest)

	conn, err := net.Dial("tcp", dest)

	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
		os.Exit(1)
	}

	return conn
}

// Update is to update the path with the sync system
func Update(path string) {
	fmt.Println(path)
	conn := connectToServer("0.0.0.0", 1248)

	// Package size can be 1 to 65495
	const BUFFERSIZE = 1024

	dat, err := ioutil.ReadFile("/home/atom/mateixTest/sample.txt")
	e.Check(err)
	text := string(dat)

	conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
	_, err = conn.Write([]byte(text))
	if err != nil {
		fmt.Println("Error writing to stream.")
	}
}
