package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
)

const (
    MAX_CONNS = 2
    BUFF_SIZE = 1 << 10
    BYE = 100
    WAIT = 101
    PLAY = 102
    END = 103
)
var (
    max_conns int
    p1r, p2r *bufio.Reader
    p1w, p2w *bufio.Writer
)

func check(tab []byte) bool {
    return tab[len(tab)] == END
}

func handle(conn net.Conn) {
    defer conn.Close()
    r := bufio.NewReader(conn)
    w := bufio.NewWriter(conn)
    buff := make([]byte, BUFF_SIZE)
    num_bytes, err := r.Read(buff)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s", err.Error())
    } else {
        if max_conns < 1 {
            fmt.Println("Player 1 connected")
            max_conns++
            p1r = r
            p1w = w
            buff[0] = WAIT
            p1w.Write(buff[:1])
            p1w.Flush()
        } else if max_conns < 2 {
            fmt.Println("Player 2 connected")
            max_conns++
            p2r = r
            p2w = w
            buff[0] = PLAY
            w.Write(buff[:1])
            w.Flush()
            gameover := false
            r = p2r
            w = p1w
            turn := 2
            for !gameover {
                num_bytes, err = r.Read(buff)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "%s", err.Error())
                }
                gameover = check(buff[:num_bytes])
                w.Write(buff[:num_bytes])
                w.Flush()
                if turn == 1 {
                    r = p1r
                    w = p2w
                    turn = 2
                } else {
                    r = p2r
                    w = p1w
                    turn = 1
                }
            }
        } else {
            fmt.Println("Player rejected")
            buff[0] = BYE
            w.Write(buff[:1])
        }
    }
}

func main() {
    ln, err := net.Listen("tcp", "0.0.0.0:8000")
    defer ln.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s", err.Error())
    } else {
        for {
            conn, err := ln.Accept()
            if err != nil {
                fmt.Fprintf(os.Stderr, "%s", err.Error())
            } else {
                go handle(conn)
            }
        }
    }
}
