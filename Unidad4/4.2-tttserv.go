package main

import (
    "bufio"
    "fmt"
    "net"
)

const (
    NEW     = byte(0) // Mensajes del cliente
    UPDATE  = byte(1)
    PLAY    = byte(2)

    WAIT    = byte(3) // Mensajes del server
    TURN    = byte(4)
)

func getMsg(buff []byte) byte {
    return buff[0]
}
func getSessId(buff []byte) byte {
    return buff[1]
}
func getPlayerId(buff []byte) byte {
    return buff[2]
}
func getTab(buff []byte) []byte {
    return buff[3:]
}
func setMsg(buff []byte, msg byte) {
    buff[0] = msg
}
func setSessId(buff []byte, sid byte) {
    buff[1] = sid
}
func setPlayerId(buff []byte, pid byte) {
    buff[2] = pid
}
func setTab(buff []byte, tab []byte) {
    for i, e := range tab {
        buff[i + 3] = e
    }
}

type Sess struct {
    turno byte
    play  bool
    tab   []byte
}

var chFirstPlayer chan bool
var currentSessId byte
var sessions map[byte]*Sess

func sesion(conn net.Conn) {
    defer conn.Close()
    r := bufio.NewReader(conn)
    w := bufio.NewWriter(conn)
    buff := make([]byte, 12)

    r.Read(buff)
    msg := getMsg(buff)
    sid := getSessId(buff)
    pid := getPlayerId(buff)
    tab := getTab(buff)
    if msg == NEW {
        if <-chFirstPlayer {
            sessions[currentSessId] = &Sess{}
            setSessId(buff, byte(currentSessId))
            setPlayerId(buff, 0)
            setMsg(buff, WAIT)
            fmt.Printf("Jugador 1 sesión %d conectado.\n", currentSessId)
            chFirstPlayer <- false
        } else {
            setSessId(buff, byte(currentSessId))
            setPlayerId(buff, 1)
            setMsg(buff, TURN)
            sessions[currentSessId].turno = 1
            fmt.Printf("Jugador 2 sesión %d conectado.\n", currentSessId)
            currentSessId++
            chFirstPlayer <- true
        }
    } else if msg == UPDATE {
        if sessions[sid].turno != pid && sessions[sid].play {
            setTab(buff, sessions[sid].tab)
            setMsg(buff, TURN)
            sessions[sid].play = false
            sessions[sid].turno = (sessions[sid].turno + 1) % 2
        } else {
            setMsg(buff, WAIT)
        }
    } else if msg == PLAY {
        sessions[sid].tab = tab
        sessions[sid].play = true
        setMsg(buff, WAIT)
    }
    w.Write(buff)
    w.Flush()
}

func main() {
    chFirstPlayer = make(chan bool, 1)
    chFirstPlayer <- true
    sessions = make(map[byte]*Sess)
    currentSessId = 1
    ln, _ := net.Listen("tcp", "localhost:8000")
    defer ln.Close()
    for {
        conn, _ := ln.Accept()
        go sesion(conn)
    }
}