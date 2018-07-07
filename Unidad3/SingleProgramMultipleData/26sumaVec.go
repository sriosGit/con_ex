package main

import "fmt"

func suma(a, b []float32) []float32 {
    if len(a) != len(b) {
        return nil
    }
    c := make([]float32, len(a))
    for i := 0; i < len(a); i++ {
        c[i] = a[i] + b[i]
    }
    return c
}

func main() {
    a := []float32 { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
    b := []float32 { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
    
    c := suma(a, b)
    
    fmt.Println(c)
}
