package main

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

func receiver(port int) {
	ServerConn, _ := net.ListenTCP("tcp", &net.TCPAddr{IP: []byte{0, 0, 0, 0}, Port: port, Zone: ""})
	defer ServerConn.Close()
	akseptert_tcp, _ := ServerConn.AcceptTCP()
	buf := make([]byte, 1024)
	reader := bytes.NewReader(buf)
	for {
		addr := ServerConn.Addr()
		n, err := akseptert_tcp.ReadFrom(reader)
		if err != nil {
			panic("panic")
		}
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
	}
}

func send(message string, port int) {
	Conn, _ := net.DialTCP("tcp", nil, &net.TCPAddr{IP: []byte{10, 100, 23, 241}, Port: port, Zone: ""})
	defer Conn.Close()
	time.Sleep(1 * time.Second)
	Conn.Write([]byte(message))
	fmt.Println("I Sent")
}

func main() {
	//go receiver(20005)
	//send("Hi Server, How are you doing?", 20005)
	go receiver(34933)

	time.Sleep(2 * time.Second)
}

//10.100.23.241:45123
