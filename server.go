package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"strconv"

	"github.com/ramantehlan/mateix/packages/command"
	"github.com/ramantehlan/mateix/packages/e"
)

// GetHash is to get the hash of a file
func GetHash(path string) string {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(path)
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

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
		command.Execute(exec.Command("sudo", "pkill", "mateix"))
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	dataFile := command.GetHome() + "/mateixTest2/data"
	conn.Write([]byte(GetHash(dataFile) + "\n"))
	scanner := bufio.NewScanner(conn)
	incomingData := ""
	for {
		ok := scanner.Scan()
		if !ok {
			break
		}
		incomingData += scanner.Text() + "\n"
	}
	if incomingData != "" {
		d1 := []byte(incomingData)
		err := ioutil.WriteFile(dataFile, d1, 0644)
		e.Check(err)
	}
	fmt.Println("Client at " + remoteAddr + " disconnected.")
}
