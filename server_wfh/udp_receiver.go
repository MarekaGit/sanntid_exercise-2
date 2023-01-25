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
		n, addr, _ := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
	}
}

func send(message string, port int) {
	Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{10, 100, 23, 241}, Port: port, Zone: ""})
	defer Conn.Close()
	Conn.Write([]byte(message))
	time.Sleep(1 * time.Second)
}

func main() {
	receiver(30000)
	send("Hi Server, How are you doing?", 20005)

}
