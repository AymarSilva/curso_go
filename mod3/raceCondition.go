package main

import (
	"fmt"
	"time"
)

func sum1(n* []int) {
	*n = append(*n,10)
	fmt.Print(*n)
}

func sum2(n* []int) {
	fmt.Print(*n)
}

func main() {
	n := []int{7, 2}

	go sum1(&n)
	go sum2(&n)

	time.Sleep(time.Second)
}

// In this code aforementioned
// I expect sum1 to print [7,2,10]
// And sum2 to print [7,2]
// But there're times when this may not be true
// Due to the race condition.
// Like sum1 and sum2 prints: [7,2,10][7,2,10]

// When programming in Go,
// there are gourotines which are managed by the go schedule,
// due to concurrency and the shared CPU between them,
// their execution got what is called interleaving,
// this means that the go scheduler decides whether
// they're going to leave goroutine1 to process goroutine2,
// this action may triggers the race condition,
// which results in:
// the system cannot define the outcomes of the program.