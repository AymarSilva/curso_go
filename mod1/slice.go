package main

import (
	"fmt"
	"sort"
	"strconv"
)

var input string
var myNum int

func main()  {
	fmt.Println("Press X for the exit")
	mySlice:= make([]int,0,3)
	for {
		fmt.Println("Enter an integer: ")
		fmt.Scan(&input)

		if input == "X"{
			break
		}

		myNum,_ = strconv.Atoi(input)
		mySlice = append(mySlice, myNum)
		sort.Ints(mySlice)
		fmt.Println("Slice: ",mySlice)
	}
}