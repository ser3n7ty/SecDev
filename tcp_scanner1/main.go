package main

import (
	"fmt"
	"net"
	"os"
	"sscanner1/utils"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: scan IP:Port ...")
		fmt.Println("\t\tscan 192.168.216.128:22,80,98 192.168.12,*:5-98 192.61.5-52:12,125")
	}

	// 存储参数
	args := os.Args[1:]
	var ipGather [][]net.IP
	var portGather [][]int

	for _, v := range args {
		parts := strings.Split(v, ":")

		ipList, err := utils.GetIPList(parts[0])
		if err != nil {
			return
		}
		portList, err := utils.GetPortList(parts[1])
		if err != nil {
			return
		}
		ipGather = append(ipGather, ipList)
		portGather = append(portGather, portList)
	}
	fmt.Println(ipGather)
	fmt.Println(portGather)

	for s := 0; s < len(ipGather); s++ {
		var ip net.IP
		var port int
		for i := 0; i < len(ipGather[s]); i++ {
			ip = ipGather[s][i]
			for j := 0; j < len(portGather[s]); j++ {
				port = portGather[s][j]
				conn, err := utils.Connect(ip, port)
				if err != nil {
					continue
				}
				fmt.Printf("IP: %s\tPort: %d \tOpen \tConn:%v\n", ip, port, conn)
			}
		}
	}

}
