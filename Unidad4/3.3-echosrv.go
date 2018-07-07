package main

import (
	"bufio"
	"net"
    "fmt"
    "strings"
)

func main() {
	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
	conn, _ := ln.Accept()
	defer conn.Close()
	r := bufio.NewReader(conn)
    for {
        // io.Copy(conn, conn)
        msg, _ := r.ReadString('\n')
        fmt.Fprintf(conn, "%s", msg)
        if strings.Contains(msg, "exit") {
            break
        }
    }
}
