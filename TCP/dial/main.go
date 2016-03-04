package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)
	// ioutil.ReadAll(conn) returns []byte, needed convert to string
	fmt.Println(string(bs))
	//fmt.Println(bs)
}
