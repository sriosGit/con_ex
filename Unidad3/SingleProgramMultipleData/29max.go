package main

import (
    "fmt"
    "time"
)

const NUMPROCS = 4

func suma(a []float32) float32 {
    max := make([]float32, NUMPROCS)
    for i := 0; i < NUMPROCS; i++ {
        go func(ini int) {
            max[ini] = a[ini]
            for k := ini; k < len(a); k += NUMPROCS {
                if a[k] > max[ini] {
                    max[ini] = a[k]
                }
            }
        }(i)
    }
    time.Sleep(time.Millisecond)
    m := max[0]
    for k := 1; k < len(max); k++ {
        if max[k] > m {
            m = max[k]
        }
    }
    
    return m
}

func main() {
    a := []float32 { 1, 2, 3, 4, 55, 6, 7, 8, 9, 10 }
    
    c := suma(a)
    
    fmt.Println(c)
}
