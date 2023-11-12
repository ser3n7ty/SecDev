package utils

import (
	"fmt"
	"github.com/malfunkt/iprange"
	"net"
	"strconv"
	"strings"
	"time"
)

func Connect(ip net.IP, port int) (net.Conn, error) {
	addr := fmt.Sprintf("%s:%d", ip.String(), port)
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)

	defer func() {
		if conn != nil {
			err := conn.Close()
			if err != nil {
				return
			}
		}
	}()

	return conn, err
}

func GetIPList(ips string) ([]net.IP, error) {
	// list 是 AddressRange 对象
	addressList, err := iprange.ParseList(ips)
	if err != nil {
		return nil, err
	}
	// list 是 ip 切片
	list := addressList.Expand()
	return list, err
}

func GetPortList(selection string) ([]int, error) {
	var ports []int
	if selection == "" {
		return ports, nil
	}
	ranges := strings.Split(selection, ",")

	for _, v := range ranges {
		if strings.Contains(v, "-") {
			parts := strings.Split(v, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid port selection segment: '%s'", v)
			}
			min, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid port number: '%v'", parts[0])
			}
			max, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid port number: '%v'", parts[1])
			}
			if min > max {
				return nil, fmt.Errorf("invalid port range: '%s'", v)
			}
			for i := min; i < max; i++ {
				ports = append(ports, i)
			}
		} else {
			if port, err := strconv.Atoi(v); err != nil {
				return nil, fmt.Errorf("invalid port number: %v", v)
			} else {
				ports = append(ports, port)
			}
		}
	}
	return ports, nil
}
