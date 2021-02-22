package main

import (
	"fmt"
	"sync"
	"time"
)

type chop_stick struct{ sync.Mutex }

type philosopher struct {
	id   int
	l, r *chop_stick
}

var eat_wg sync.WaitGroup

func main() {

	count := 5

	Csticks := make([]*chop_stick, count)
	for i := 0; i < count; i++ {
		Csticks[i] = new(chop_stick)
	}

	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, l: Csticks[i], r: Csticks[(i+1)%count]}
		eat_wg.Add(1)
		go philosophers[i].eat()

	}
	eat_wg.Wait()
}

func (p philosopher) eat() {
	for j := 0; j < 3; j++ {
		p.l.Lock()
		p.r.Lock()

		fmt.Printf("phi %d is eating\n", p.id+1)
		time.Sleep(time.Second)

		p.r.Unlock()
		p.l.Unlock()

		fmt.Printf("phi %d is finished eating\n", p.id+1)
		time.Sleep(time.Second)
	}
	eat_wg.Done()
}
