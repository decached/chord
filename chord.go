package main

import (
	"crypto/sha256"
	"encoding/binary"
	"net"
	"net/http"
	"net/rpc"
)

type Node struct {
	ID    uint64
	IP    string
	Port  string
	Table map[uint64]string
}

func NewNode(ip string, port string) (node *Node, e error) {
	id := hash(ip + ":" + port)
	node = &Node{
		ID:    id,
		IP:    ip,
		Port:  port,
		Table: make(map[uint64]string),
	}
	e = node.listen(port)
	if e != nil {
		return nil, e
	}
	return node, e
}

func (n *Node) listen(port string) error {
	rpc.Register(n)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+port)
	if e != nil {
		return e
	}
	go http.Serve(l, nil)
	return nil
}

func (n *Node) Get(k string, val *string) error {
	key := hash(k)
	if v, ok := n.Table[key]; ok {
		*val = v
		return nil
	}
	// FIXME: Find `key` on other nodes.
	return nil
}

func hash(k string) uint64 {
	h := sha256.New()
	h.Write([]byte(k))
	return binary.BigEndian.Uint64(h.Sum(nil))
}
