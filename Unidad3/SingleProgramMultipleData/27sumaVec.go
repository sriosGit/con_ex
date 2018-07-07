package main

import "fmt"

const NUMPROCS = 4

func suma(a, b []float32) []float32 {
    if len(a) != len(b) {
        return nil
    }
    c := make([]float32, len(a))
    n := len(a)
    s := n / NUMPROCS
    if n % NUMPROCS != 0 {
        s += 1
    }
    for i := 0; i < NUMPROCS; i++ {
        go func(ini, fin int) {
            for k := ini; k < fin || k < len(a); k++ {
                c[k] = a[k] + b[k]
            }
        }(i * s, (i+1) * s)
    }
    return c
}

func main() {
    a := []float32 { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
    b := []float32 { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
    
    c := suma(a, b)
    
    fmt.Println(c)
}
