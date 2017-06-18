package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":42425")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		res, _ := ioutil.ReadAll(conn)
		fmt.Println(string(res))
		conn.Write(res)
	}
}
