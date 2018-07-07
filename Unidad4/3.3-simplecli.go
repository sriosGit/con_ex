package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	r := bufio.NewReader(conn)
	fmt.Fprint(conn, "Rosa\n")
	msg, _ := r.ReadString('\n')
	fmt.Println(msg)
}