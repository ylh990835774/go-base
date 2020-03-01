package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type ScanSafeCount struct {
	count int
	mux   sync.Mutex
}

var scanCount ScanSafeCount

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	inport := make(chan int)
	outport := make(chan int)
	var collect []int

	if len(os.Args) != 4 {
		fmt.Println("使用形式: port_scanner IP startPort endPort")
		os.Exit(0)
	}

	sTime := time.Now().Unix()

	ip := string(os.Args[1])
	startPort, _ := strconv.Atoi(os.Args[2])
	endPort, _ := strconv.Atoi(os.Args[3])

	if startPort > endPort {
		fmt.Println("Usage: scanner IP startPort endPort")
		fmt.Println("endPort must be larger than startPort")
		os.Exit(0)
	}
	scanCount = ScanSafeCount{count: (endPort - startPort + 1)}

	fmt.Printf("扫描 %s: %d-------%d\n", ip, startPort, endPort)

	go loop(inport, startPort, endPort)

	for v := range inport {
		go scanner(v, outport, ip, endPort)
	}

	for port := range outport {
		if port != 0 {
			collect = append(collect, port)
		}
	}

	fmt.Println("--")
	fmt.Println(collect)
	eTime := time.Now().Unix()
	fmt.Println("扫描时间:", eTime-sTime)
}

func scanner(inport int, outport chan int, ip string, endPort int) {
	in := inport

	host := fmt.Sprintf("%s:%d", ip, in)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		outport <- 0
	} else {
		conn, err := net.DialTimeout("tcp", tcpAddr.String(), 1*time.Second)
		if err != nil {
			outport <- 0
		} else {
			outport <- in
			fmt.Printf("\n**********（%d 可以）********\n", in)
			_ = conn.Close()
		}
	}

	scanCount.mux.Lock()
	scanCount.count = scanCount.count - 1
	if scanCount.count <= 0 {
		close(outport)
	}
	scanCount.mux.Unlock()
}

func loop(inport chan int, startPort int, endPort int) {
	for i := startPort; i < endPort; i++ {
		inport <- i
	}
	close(inport)
}
