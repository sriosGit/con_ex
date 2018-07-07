package main

import (
    "fmt"
    "time"
)

const NUMPROCS = 4

func suma(a, b []float32) []float32 {
    if len(a) != len(b) {
        return nil
    }
    c := make([]float32, len(a))
    for i := 0; i < NUMPROCS; i++ {
        go func(ini int) {
            for k := ini; k < len(a); k += NUMPROCS {
                c[k] = a[k] + b[k]
            }
        }(i)
    }
    time.Sleep(time.Millisecond)
    return c
}

func main() {
    a := []float32 { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
    b := []float32 { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
    
    c := suma(a, b)
    
    fmt.Println(c)
}
