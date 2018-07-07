package main

import (
    "fmt"
)

func cualquierCosa(fin chan bool) {
    for i := 0; i < 10; i++ {
        fmt.Printf("Hola %d\n", i)
    }
    fin <- true
} 

func main() {
    fin := make(chan bool)
    go cualquierCosa(fin)
    <- fin
}