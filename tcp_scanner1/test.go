package main

import (
	"fmt"
	"github.com/malfunkt/iprange"
	"log"
)

func main() {
	list, err := iprange.ParseList("10.0.0.1, 10.0.0.5-10, 192.168.1.*, 192.168.10.0/24")
	if err != nil {
		log.Printf("error: %s", err)
	}
	// list 是 AddressRange 对象
	fmt.Println(list)
	// rng 是 ip 切片
	rng := list.Expand()
	fmt.Println(rng)
}
