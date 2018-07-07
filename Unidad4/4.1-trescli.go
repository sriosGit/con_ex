package main

import (
    "fmt"
    "net"
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

func verify(tab []byte, row, col int, p byte) bool {
    if tab[row * 3 + col] == 0 {
        tab[row * 3 + col] = p
        return true
    }
    return false
}
func check(tab []byte) bool {
    if (tab[0] == tab[1] && tab[1] == tab[2]) ||
        (tab[3] == tab[4] && tab[4] == tab[5]) ||
        (tab[6] == tab[7] && tab[7] == tab[8]) ||
        (tab[0] == tab[3] && tab[3] == tab[6]) ||
        (tab[1] == tab[4] && tab[4] == tab[7]) ||
        (tab[2] == tab[5] && tab[5] == tab[8]) ||
        (tab[0] == tab[4] && tab[4] == tab[8]) ||
        (tab[2] == tab[4] && tab[4] == tab[6]) {
        return true
    }
    for i := 0; i < 9; i++ {
        if tab[i] == 0 {
            return false
        }
    }
    return true
}

func piece(tab []byte) byte {
    for i := 0; i < 9; i++ {
        if tab[i] == 1 {
            return 2
        } else if tab[i] == 2 {
            return 1
        }
    }
    return 1
}

func main() {
    conn, _ := net.Dial("tcp", "0.0.0.0:8000")
    defer conn.Close()
    r := bufio.NewReader(conn)
    w := bufio.NewWriter(conn)
    buff := make([]byte, BUFF_SIZE)
    w.Write(buff[:1])
    w.Flush()
    r.Read(buff)
    var p byte
    gameover := false
    if buff[0] == WAIT {
        n, _ = r.Read(buff)
        fmt.Printf("%d\n", n)
        p = piece(buff)
    } else {
        fmt.Print("Ingresa ficha [1: X, 2: O]: ")
        fmt.Scanf("%d", &p)
    }
    var (
        row, col int
        valid bool
    )
    for !gameover {
        valid = false
        for valid {
            fmt.Print("Ingrese jugada [fila columna]: ")
            fmt.Scanf("%d %d", &row, &col)
            valid = verify(buff[:9], row, col, p)
        }
        gameover = check(buff)
        w.Write(buff[:10])
        w.Flush()
        if !gameover {
            r.Read(buff)
            gameover = check(buff)
        }
    }
}
