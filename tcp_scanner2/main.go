package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"tcp_scanner2/scanner"
	"tcp_scanner2/utils"
)

// 并发需要达到的效果是 一个协程一次性处理一个 IP:Port 的 Connect() 测试
func main() {

	if len(os.Args) < 2 {
		_ = fmt.Errorf("usage: <PROGRAM_NAME> IP:Port ")
		return
	}

	args := os.Args[1:]

	var ipGather [][]net.IP
	var portGather [][]int
	// 声明 map
	var tasks []map[string]int

	for _, v := range args {
		parts := strings.Split(v, ":")
		fmt.Println(parts)
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

	tasks = scanner.GenerateTask(ipGather, portGather)
	fmt.Println(tasks)
	scanner.AssignTask(tasks)
	scanner.PrintResult()
}

// 包被导入时自动执行，不需要显式调用
// runtime.NumCPU():获取当前计算机上可用的 CPU 核心数量
// runtime.GOMAXPROCS(): 用于设置 GO 程序的最大并发执行线程数
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
