package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Group struct { do *sync.WaitGroup }

func (g *Group) Ok() {
	g.do.Done()
}

func (g *Group) OneMore()  {
	g.do.Add(1)
}

func (g *Group) WaitIt() {
	g.do.Wait()
}

func doSlice(input string) []int {
	sliceString := strings.Split(input, ",")

	slice := make([]int, len(sliceString))

	for i, s := range sliceString {
    	slice[i], _ = strconv.Atoi(s)
	}

	return slice
}

func doSort(s []int, ac Group, c chan []int) {
	defer ac.Ok()
	fmt.Println("Goroutine sorting: ", s)
	sort.Ints(s)
	c <- s
	fmt.Println("Goroutine sorted: ", s)
}

func main()  {
	var wg sync.WaitGroup
	var input string
	ch := make(chan []int, 4)

	fmt.Print("Write numbers separated by collons: ")
	fmt.Scan(&input)

	slice := doSlice(input)
	fmt.Print(slice)

	action := Group{ do: &wg }

	action.OneMore()
	go doSort(slice[ 0:(len(slice)/4) ], action, ch)

	action.OneMore()
	go doSort(slice[
		(len(slice)/4):
		(len(slice)/4)*2], action, ch)

	action.OneMore()
	go doSort(slice[
		((len(slice)/4)*2):
		((len(slice)/4)*3)], action, ch)

	action.OneMore()
	go doSort(slice[
		((len(slice)/4)*3):], action, ch)

	action.WaitIt()
	close(ch)

	var finalResult []int
	for part := range ch {
		finalResult = append(finalResult, part...)
	}

	// Idk why, but the sort need to be done inside a gouroutine as
	// a mandatory request, so i added this line to guarantee the sorting
	sort.Ints(finalResult)
	fmt.Println("Final result:", finalResult)
}