package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func packetSlitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == 0x123456 {
		var l int16
		_ = binary.Read(bytes.NewBuffer(data[4:6]), binary.BigEndian, &l)
		pl := int(l) + 6
		if pl < len(data) {
			return pl, data[:pl], nil
		}
	}
	return
}

func main() {
	l, err := net.Listen("tcp", ":4044")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen to 4044")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		} else {
			//go handleConn(conn)
			go handleConn2(conn)
		}
	}
}

// 解决粘包问题的handle
func handleConn2(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("关闭")

	fmt.Println("新连接：", conn.RemoteAddr())

	result := bytes.NewBuffer(nil)
	var buf [65542]byte

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read err:", err)
				break
			}
		} else {
			scanner := bufio.NewScanner(result)
			scanner.Split(packetSlitFunc)
			for scanner.Scan() {
				fmt.Println("recv:", string(scanner.Bytes()[6:]))
			}
		}
		result.Reset()
	}
}

// 会产生粘包问题的处理handle
func handleConn(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("关闭")

	result := bytes.NewBuffer(nil)
	var buf [1024]byte

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read err:", err)
				break
			}
		} else {
			fmt.Println("recv:", result.String())
		}
		result.Reset()
	}
}
