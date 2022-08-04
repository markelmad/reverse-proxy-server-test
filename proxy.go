package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "8-f48cf3a682-7fthvk.api.dev.zesty.io"
	CONN_PORT = ":8080"
	CONN_TYPE = "tcp"
)

func proxy() {

	l, err := net.Listen(CONN_TYPE, CONN_HOST+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	conn.Write([]byte("Message received."))
	fmt.Println("Message received!!!!.")

	conn.Close()
}
