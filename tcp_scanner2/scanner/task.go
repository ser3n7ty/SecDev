package scanner

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"tcp_scanner2/vars"
)

func GenerateTask(ipGather [][]net.IP, portGather [][]int) []map[string]int {
	var tasks []map[string]int
	// ipGather 和 portGather 在纵向维度是等量的
	for s := 0; s < len(ipGather); s++ {
		for i := 0; i < len(ipGather[s]); i++ {
			for j := 0; j < len(portGather[s]); j++ {
				ip := ipGather[s][i].String()
				port := portGather[s][j]
				pair := map[string]int{ip: port}
				tasks = append(tasks, pair)
			}
		}
	}
	return tasks
}

func AssignTask(tasks []map[string]int) {
	// 需要进行的轮询次数
	scanBatch := len(tasks) / vars.ThreadNum
	fmt.Println(scanBatch)
	for i := 0; i < scanBatch; i++ {
		task := tasks[vars.ThreadNum*i : vars.ThreadNum*(1+i)]
		fmt.Println(task)
		RunTask(task)
	}
	if len(tasks)%vars.ThreadNum > 0 {
		lastTask := tasks[scanBatch*vars.ThreadNum:]
		fmt.Println(lastTask)
		RunTask(lastTask)
	}
}
func RunTask(tasks []map[string]int) {
	var wg sync.WaitGroup
	// 将 WaitGroup 的计数器设置为任务的数量
	wg.Add(len(tasks))
	for _, task := range tasks {
		for ip, port := range task {
			go func(ip string, port int) {
				SaveResult(Connect(ip, port))
				// 减少计数器，表示一个任务已完成
				wg.Done()
			}(ip, port)
		}
	}
	// 阻塞主进程
	wg.Wait()
}

func SaveResult(ip string, port int, err error) {
	if err != nil {
		return
	}
	// 只保存端口开放的记录
	// v 存储指定键对应的值的变量
	// ok 表示 Map 中是否存储了 ip
	// v 是 ip 对应的键，类型是空接口 interface{}
	v, isExist := vars.Result.Load(ip)
	if isExist {
		// 类型断言
		ports, ok := v.([]int)
		if ok {
			ports = append(ports, port)
			vars.Result.Store(ip, ports)
		}
	} else {
		ports := make([]int, 0)
		ports = append(ports, port)
		vars.Result.Store(ip, ports)
	}
}

func PrintResult() {
	// 函数原型：结构体函数
	// func (m *Map) Range(f func(key, value any) bool)
	vars.Result.Range(func(key, value interface{}) bool {
		fmt.Printf("IP: %v\n", key)
		fmt.Printf("Ports: %v\n", value)
		fmt.Println(strings.Repeat("-", 100))
		return true
	})
}
