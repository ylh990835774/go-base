package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	data := []byte("[这是一个完整的数据包]123456")
	conn, err := net.DialTimeout("tcp", "localhost:4044", 30*time.Second)
	if err != nil {
		fmt.Printf("connect failed, err:%v\n", err)
		return
	}
	for i := 0; i < 1000; i++ {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write failed, err:%v\n", err)
			break
		}
	}
}
