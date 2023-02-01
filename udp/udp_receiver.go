package main

import (
	"fmt"
	"net"
	"time"
)

func receiver(port int) {
	ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: port, Zone: ""})
	defer ServerConn.Close()
	buf := make([]byte, 1024)
	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			panic("panic")
		}
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
	}
}

func send(message string, port int) {
	Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{10, 100, 23, 241}, Port: port, Zone: ""})
	defer Conn.Close()
	time.Sleep(1 * time.Second)
	Conn.Write([]byte(message))
	fmt.Println("I Sent")
}

func main() {
	go receiver(20005)
	send("Hi Server, How are you doing?", 20005)

	time.Sleep(2 * time.Second)
}

//10.100.23.241:45123
