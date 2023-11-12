package scanner

import (
	"fmt"
	"net"
	"time"
)

func Connect(ip string, port int) (string, int, error) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	// 记得关闭连接
	defer func() {
		if conn != nil {
			err := conn.Close()
			if err != nil {
				return
			}
		}
	}()
	fmt.Println(port, ip)
	return ip, port, err
}
