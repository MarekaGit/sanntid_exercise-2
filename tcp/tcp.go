package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	HOST = "10.100.23.11"
	PORT = "34933"
	TYPE = "tcp"
)

func receiver(port string) {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+port)

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
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println("Received message:", string(received))

	conn.Close()
}

func send(message string, port string) {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+port)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte("This is a message"))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println("I Sent")
	conn.Close()
}

func main() {
	go receiver("20005")
	send("Hi Server, How are you doing?", "20005")
	//go receiver(34933)

	time.Sleep(2 * time.Second)
}

//10.100.23.11:59300

/*package main

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
		n, err := akseptert_tcp.Read(reader)
		if err != nil {
			panic("panic")
		}
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
	}
}

func send(message string, port int) {
	Conn, _ := net.DialTCP("tcp", nil, &net.TCPAddr{IP: []byte{10, 100, 23, 11}, Port: port, Zone: ""})
	defer Conn.Close()
	time.Sleep(1 * time.Second)
	Conn.Write([]byte(message))
	fmt.Println("I Sent")
}

func main() {
	go receiver(20005)
	send("Hi Server, How are you doing?", 20005)
	//go receiver(34933)

	time.Sleep(2 * time.Second)
}*/

//10.100.23.11:59300
