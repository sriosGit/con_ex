package main

import (
    "fmt"
)

func main() {
    ch := make(chan int, 1)
    ch <- 5
    x := <- ch
    fmt.Printf("Valor: %d\n", x)
}