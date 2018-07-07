package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const N_PROCESS = 100

func main() {
	//array := []int{9, 6, 7, 3, 2, 4, 5, 1}
	jobs := make(chan int, N_PROCESS)
	results := make(chan []int)

	go insertion(jobs, results)

	for i := 0; i < N_PROCESS; i++ {
		jobs <- rand.Intn(100)
	}
	close(jobs)
	fmt.Println(<-results)
}
func insertion(jobs <-chan int, results chan<- []int) {
	var array []int
	for i := 1; i < N_PROCESS; i++ {
		for n := range jobs {
			array = append(array, n)
		}
	}
	sort.Ints(array)
	results <- array

}
