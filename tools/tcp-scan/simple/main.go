package simple

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

// 定义可接受的参数
var (
	// 需要检测的域名
	hostname string
	// 开始端口
	startPort int
	// 结束端口
	endPort int
	// 超时时间
	timeout time.Duration
)

func init() {
	flag.StringVar(&hostname, "hostname", "www.baidu.com", "hostname to test")
	flag.IntVar(&startPort, "start-port", 80, "the port on which the scanning starts")
	flag.IntVar(&endPort, "end-port", 100, "the port on which the scanning ends")
	flag.DurationVar(&timeout, "timeout", time.Millisecond*200, "timeout")
}

func main() {
	flag.Parse()

	// 保存可以访问的端口数组
	var ports []int

	// 引入并发协程
	wg := &sync.WaitGroup{}
	// 引入互斥锁，保证写入ports的数据在多协程中是正确的
	mutex := &sync.Mutex{}

	for port := startPort; port < endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(hostname, p, timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("\nopened ports: %v\n", ports)
}

// 检测一个域名及端口是否打开
func isOpen(host string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp",
		fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		fmt.Printf("connect [%s:%d] error: %v\n", hostname, port, err)
		return false
	}
	_ = conn.Close()
	return true
}
