package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"

	"github.com/ramantehlan/mateix/packages/command"
	"github.com/ramantehlan/mateix/packages/e"
)

func listenToClinet(addr string, port int) net.Listener {
	src := addr + ":" + strconv.Itoa(port)
	server, err := net.Listen("tcp", src)
	e.Check(err)
	fmt.Printf("Listening on %s.\n", src)
	return server
}

func server(action bool) {
	if action {
		// Start the server
		fmt.Println("Starting server")
		server := listenToClinet("0.0.0.0", 1248)
		defer server.Close()

		for {
			conn, err := server.Accept()
			e.Check(err)
			go handleConnection(conn)
		}

	} else {
		// End the server
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	scanner := bufio.NewScanner(conn)
	file := ""
	for {
		ok := scanner.Scan()
		if !ok {
			break
		}
		file += scanner.Text() + "\n"
	}
	d1 := []byte(file)
	err := ioutil.WriteFile(command.GetHome()+"/mateixTest/sample2.txt", d1, 0644)
	e.Check(err)
	fmt.Println("Client at " + remoteAddr + " disconnected.")
}
