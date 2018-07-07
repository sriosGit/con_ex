package main

import (
    "fmt"
    "time"
    "net"
    "bufio"
)

const (
	NEW    = byte(0) // Mensajes del cliente
	UPDATE = byte(1)
	PLAY = byte(2)

	WAIT  = byte(3) // Mensajes del server
	TURN = byte(4)
)

func printTab(tab []byte) {
	fmt.Println("+---+---+---+")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("| %c ", tab[i*3 + j])
		}
		fmt.Println("|\n+---+---+---+")
	}
}

func scanMove(tab []byte, p byte) {
	var i, j byte
	valid := false
	for !valid {
		fmt.Printf("Jugada para %c [0-2] [0-2]: ", rune(p))
		fmt.Scanf("%d %d\n", &i, &j)
		idx := i * 3 + j
        if i >= 0 && i < 3 && j >= 0 && j < 3 && tab[idx] == byte(' ') {
			tab[idx] = p
			valid = true
		} else {
			fmt.Println(" --- Jugada no permitida ---")
		}
	}
}

func findWinner(tab []byte) byte {
	for i := 0; i < 3; i++ {
        if (tab[i*3] == tab[i*3+1] && tab[i*3+1] == tab[i*3+2]) ||
           (tab[i] == tab[i+3] && tab[i+3] == tab[i+6]) {
            if tab[i*4] != byte(' ') {
                return tab[i*4]
            }
		}
	}
	if tab[4] != byte(' ') &&
       ((tab[0] == tab[4] && tab[4] == tab[8]) ||
        (tab[2] == tab[4] && tab[4] == tab[6])) {
	   	return tab[4]
	}
	for i := 0; i < 9; i++ {
		if tab[i] != 0 {
			return 0;
		}
	}
	return byte('-') // empate
}

func chooseOpositeToken(tab [] byte) byte {
	var token rune
	for i := 0; i < 9; i++ {
		if tab[i] != byte(' ') {
			token = rune(tab[i])
			break
		}
	}
	if token == 'o' {
		return byte('x')
	} else if token == 'x' {
		return byte('o')
	}
	return 0
}

func pickToken() byte {
	p := '-'
	for p != 'o' && p != 'x' {
		fmt.Println("Seleccione ficha [x,o]: ");
		fmt.Scanf("%c\n", &p)
		if p != 'o' && p != 'x' {
			fmt.Println(" --- Ficha no permitida ---")
		}
	}
	return byte(p)
}

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

func main() {
    var piece byte
    gameover := false
    buff := []byte { NEW, 0, 0, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ' }
    starting := true
    var winner byte = 0
    for !gameover {
        conn, _ := net.Dial("tcp", "localhost:8000")
        defer conn.Close()
        r := bufio.NewReader(conn)
        w := bufio.NewWriter(conn)
        
        w.Write(buff)
        w.Flush()
        
        if winner != 0 {
            break
        }
                
        r.Read(buff)
        msg := getMsg(buff)
        tab := getTab(buff)
        
        winner = findWinner(tab)
        
        if winner == 0 {
            if msg == WAIT {
                setMsg(buff, UPDATE)
            } else {
                if starting {
                    starting = false
                    piece = chooseOpositeToken(tab)
                    if piece == 0 {
                        piece = pickToken()
                    }
                }
                printTab(tab)
                scanMove(tab, piece)
                printTab(tab)
                setMsg(buff, PLAY)
                setTab(buff, tab)
                winner = findWinner(tab)
            }
            time.Sleep(time.Second)
        } else {
            gameover = true
        }
    }
    fmt.Printf("Ganador %c\n", rune(winner))
}