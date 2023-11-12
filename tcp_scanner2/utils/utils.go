package utils

import (
	"fmt"
	"github.com/malfunkt/iprange"
	"net"
	"strconv"
	"strings"
)

func GetIPList(ips string) ([]net.IP, error) {
	addressList, err := iprange.ParseList(ips)
	if err != nil {
		return nil, err
	}
	ipList := addressList.Expand()
	return ipList, err
}
func GetPortList(selection string) ([]int, error) {
	var ports []int
	if len(selection) == 0 {
		return nil, nil
	}

	ranges := strings.Split(selection, ",")
	fmt.Printf("ranges: %v\n", ranges)
	for _, v := range ranges {
		if strings.Contains(v, "-") {
			parts := strings.Split(v, "-")
			fmt.Println(parts)
			min, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid port number: '%s'", parts[0])
			}
			max, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid port number: '%s'", parts[1])
			}

			if max < min {
				return nil, fmt.Errorf("invalid port range: '%s'", v)
			}
			for i := min; i < max; i++ {
				ports = append(ports, i)
			}
			fmt.Printf("ports: %v\n", ports)

		} else {
			if port, err := strconv.Atoi(v); err != nil {
				return nil, fmt.Errorf("invalid port number: '%s'", v)
			} else {
				ports = append(ports, port)
			}
		}
	}

	return ports, nil

}
