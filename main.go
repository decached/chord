package main

import (
	"fmt"
	"os"
)

func main() {
	node, e := NewNode(os.Args[1], os.Args[2])
	if e != nil {
		panic(e)
	}
	fmt.Printf("Node running on %s:%s\n", node.IP, node.Port)
	for {
		var cmd string
		fmt.Printf("> ")
		fmt.Scanf("%s", &cmd)
		switch cmd {
		case "get":
			var key, val string
			fmt.Scanf("%s", &key)
			e = node.Get(key, &val)
			fmt.Printf("%s\n", val)
		}
	}
}
