package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	Name string
}

func (p *Philosopher) Eat(left, right chan bool) {
	for {
	Left:
		select {
		case <-left:
			for {
				select {
				case <-right:
					fmt.Println(p.Name, "is eating!")
					time.Sleep(time.Second)
					left <- true
					right <- true
					break Left
				default:
					left <- true
					break Left
				}
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func eat(p Philosopher, left, right chan bool) {
	go p.Eat(left, right)
}

func main() {
	p := []Philosopher{
		Philosopher{"A"},
		Philosopher{"B"},
		Philosopher{"C"},
		Philosopher{"D"},
		Philosopher{"E"},
	}

	forks := []chan bool{}
	for range p {
		forks = append(forks, make(chan bool, 1))
	}

	for i, pp := range p {
		eat(pp, forks[i], forks[(i+1)%len(forks)])
	}

	for i := range forks {
		forks[i] <- true
	}

	time.Sleep(10 * time.Second)
}
