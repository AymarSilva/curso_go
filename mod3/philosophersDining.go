package main

import (
	"fmt"
	"sync"
)

type ChopS struct { sync.Mutex }

type Philo struct {
	rightS *ChopS
	leftS *ChopS
	number int
}

func (p *Philo) eat(w *sync.WaitGroup) {
	if (p.number != 4){ // Due to interleaving, philo 4 may trigger a deadlock so it eats reversely
		p.leftS.Lock()
		p.rightS.Lock()
	}else{
		p.rightS.Lock()
		p.leftS.Lock()
	}

	fmt.Print("Starting to eat ",p.number,"\n")

	fmt.Print("Finishing to eat ",p.number,"\n")

	p.leftS.Unlock()
	p.rightS.Unlock()
	w.Done()
}

func main() {
	var w sync.WaitGroup // In order to wait the concurrent void method do it's own

	var cs = make([]*ChopS, 5)
	for i:= 0; i < len(cs); i++ {
		cs[i] = &ChopS{}
	}

	var philos = make([]*Philo, 5) // Slice of five philos of chopS
	for i:= 0; i < len(philos); i++ {
		philos[i] = &Philo{
			number: i+1,
			leftS: cs[i],
			rightS: cs[ (i+1) %5 ],
		}
	}

	for i := 0; i < 3; i++ { // Each philosopher must run three times
		for j := 0; j < len(philos); j++ {
			w.Add(1)
			go philos[j].eat(&w)
		}
	}
	w.Wait()
}