package main

import (
	"crypto/sha256"
	"encoding/binary"
)

type Node struct {
	ID    uint64
	Table map[uint64]string
}

func hash(k string) uint64 {
	h := sha256.New()
	h.Write([]byte(k))
	return binary.BigEndian.Uint64(h.Sum(nil))
}

func (n *Node) Get(k string) string {
	key := hash(k)
	if v, ok := n.Table[key]; ok {
		return v
	}
	// FIXME: Find `key` on other nodes.
	return ""
}

func (n *Node) PutRandom() {
	// FIXME: Put random `key:values` in n.Table
}

func NewNode(host string, port string) (node *Node, e error) {
	id := hash(host + ":" + port)
	node = &Node{
		ID:    id,
		Table: make(map[uint64]string),
	}
	return node, e
}
