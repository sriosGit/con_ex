// Santa
package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	ID     = 0
	STATUS = 1
)

// possible states and commands
const (
	STOP  = 0
	WORK  = 1
	DONE  = 2
	READY = 3
)

const (
	NUMBER_OF_REINDEERS = 9
	NUMBER_OF_ELVES     = 10
	NEEDED_EVLES        = 3
	NUMBER_OF_RUNS      = 100
)

func main() {
	var wg sync.WaitGroup

	sendReindeers := make([]chan int, NUMBER_OF_REINDEERS)
	sendElves := make([]chan int, NUMBER_OF_ELVES)

	receiveReindeers := make(chan []int)
	receiveElves := make(chan []int)

	for i := range sendReindeers {
		sendReindeers[i] = make(chan int)

		wg.Add(1)
		go func(i int) {
			Reindeer(i, receiveReindeers, sendReindeers[i])
			wg.Done()
		}(i)
	}

	for j := range sendElves {
		sendElves[j] = make(chan int)

		wg.Add(1)
		go func(j int) {
			Elf(j, receiveElves, sendElves[j])
			wg.Done()
		}(j)
	}

	wg.Add(1)
	go func() {
		Santa(receiveReindeers, sendReindeers, receiveElves, sendElves)
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("Main ended\n")
}

func Santa(rReceive chan []int, rSend []chan int, eReceive chan []int, eSend []chan int) {
	fmt.Printf("Santas is doing some work...\n")

	var rCnt int
	var chanDeers []chan int
	resetReindeers := func() {
		chanDeers = make([]chan int, NUMBER_OF_REINDEERS)
		rCnt = 0
	}
	resetReindeers()

	var eCnt int
	var chanElves []chan int
	resetElves := func() {
		chanElves = make([]chan int, NEEDED_EVLES)
		eCnt = 0
	}
	resetElves()

	wait := func(helperReceive chan int) {
		for {
			select {
			case status := <-helperReceive:
				switch status {
				case DONE:
					return
				default:
				}
			}
		}
	}

	for a := 0; a <= NUMBER_OF_RUNS; a++ {
		fmt.Println("NUMBER_OF_RUNS DONE = ", a)
		select {
		case rArr := <-rReceive:
			if rStatus := rArr[STATUS]; rStatus == READY {
				chanDeers[rArr[ID]] = rSend[rArr[ID]]
				rCnt++

				if rCnt == NUMBER_OF_REINDEERS {
					fmt.Printf("ALL REINDEERS ARE BACK FROM HOLIDAY\n")
					for _, chanDeer := range chanDeers {
						chanDeer <- WORK
						wait(chanDeer)
					}
					resetReindeers()
				}
			}
		case eArr := <-eReceive:
			if eStatus := eArr[STATUS]; eStatus == READY {
				chanElves[eCnt] = eSend[eArr[ID]]
				eCnt++
				select {
				case deArr := <-rReceive:
					if deStatus := deArr[STATUS]; deStatus == READY {
						chanDeers[deArr[ID]] = rSend[deArr[ID]]
						rCnt++

						if rCnt == NUMBER_OF_REINDEERS {
							fmt.Printf("ALL REINDEERS ARE BACK FROM HOLIDAY\n")
							for _, chanDeer := range chanDeers {
								chanDeer <- WORK
								wait(chanDeer)
							}
							resetReindeers()
						}
					}
				default:
				}
				if eCnt == NEEDED_EVLES {
					fmt.Println("CONSULT WITH ELVES ABOUT R&D")
					for _, chanElf := range chanElves {
						chanElf <- WORK
						wait(chanElf)
					}
					resetElves()
				}
			}
		case <-time.After(time.Second * 3):
			fmt.Printf("Timed out!")
			break
		}
	}

	for _, s := range rSend {
		s <- STOP
	}

	for _, s := range eSend {
		s <- STOP
	}

	for _, s := range rSend {
		close(s)
	}

	for _, s := range eSend {
		close(s)
	}
}

func Reindeer(id int, contact chan<- []int, communicate chan int) {
	for {
		select {
		case contact <- []int{id, READY}:
			fmt.Printf("Deer %d is back from holiday\n", id)
			select {
			case w := <-communicate:
				switch w {
				case WORK:
					fmt.Printf("Deer %d is delivering presents\n", id)
					communicate <- DONE
					time.Sleep(500 * time.Millisecond)
				}
			}
		case cmd := <-communicate:
			switch cmd {
			case STOP:
				fmt.Printf("Deer %d work is done...\n", id)
				return
			}
		}
	}
}

func Elf(id int, contact chan<- []int, communicate chan int) {
	for {
		select {
		case contact <- []int{id, READY}:
			fmt.Printf("Elf %d is ready to do some work\n", id)
			select {
			case w := <-communicate:
				switch w {
				case WORK:
					fmt.Printf("Elf %d is making presents\n", id)
					communicate <- DONE
					time.Sleep(500 * time.Millisecond)
				}
			}
		case cmd := <-communicate:
			switch cmd {
			case STOP:
				fmt.Printf("Elf %d work is done...\n", id)
				return
			}
		}
	}
}
