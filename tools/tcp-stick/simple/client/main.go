package main

import (
	"fmt"
	"net"
	"os"
	"protocol"
	"time"
)

func main() {
	server := "127.0.0.1:9988"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Println("connect success")
	go sender(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}

func sender(conn *net.TCPConn) {
	for i := 0; i < 1000; i++ {
		words := fmt.Sprintf("{\"Id\":%d,\"Name\":\"golang\",\"Message\":\"message\"}", i)
		_, _ = conn.Write(protocol.Packet([]byte(words)))
	}
	fmt.Println("send over")
}
