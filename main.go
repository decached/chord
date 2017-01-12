package main

import (
	"fmt"
	"os"
)

func main() {
	node, e := NewNode("localhost", "9090")
	if e != nil {
		os.Exit(1)
	}
}
