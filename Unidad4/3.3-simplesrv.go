package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
	conn, _ := ln.Accept()
	defer conn.Close()
	r := bufio.NewReader(conn)
	name, _ := r.ReadString('\n')
	fmt.Fprintf(conn, "Hola, %s\n", name)
}
