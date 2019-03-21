package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func server(action bool) {
	fmt.Println(action)
	if action {
		ssh.Handle(func(s ssh.Session) {
			io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
		})

		log.Println("starting ssh server on port 1248")
		log.Fatal(ssh.ListenAndServe(":1248", nil))

	} else {
		// Code to stop the server
	}
}
