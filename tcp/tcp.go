package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	HOST = "10.100.23.11"
	//PORT = "33546"
	TYPE = "tcp"
)

func receive_msg(PORT string) {
	fmt.Println("hei")
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	// buffer to get data
	received := make([]byte, 1024)
	for {
		_, err = conn.Read(received)
		if err != nil {
			println("Read data failed:", err.Error())
			os.Exit(1)
		}
		println("Received message:", string(received))
	}

	//conn.Close()
}

func send_msg(PORT string, msg string) {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write([]byte(msg))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println("I sent a message")
	conn.Close()
}

func main() {
	send_msg("34933", "Heisveis")
	go receive_msg("34933")

	time.Sleep(2 * time.Second)
}

