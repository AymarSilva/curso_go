package main

import(
	"fmt"
	"strconv"
)

func BubbleSort(slice []int)  {
	for i:= 0; i < len(slice); i++ {
		for j:= 0; j < (len(slice)-1); j++ {
			swap(slice, j)
		}
	}
}

func swap(slice []int, index int)  {
	x := slice[index]
	y := slice[index+1]
	
	if x > y {
		slice[index] = y
		slice[index+1] = x
	}
}

func main() {
	sl:= make([]int, 0, 10)

	for {
		var input string
		fmt.Scan(&input)

		if input == "stop" {
			fmt.Println("Exiting...")
			break
		}

		number,_ := strconv.Atoi(input)
		sl = append(sl, number)
	}

	BubbleSort(sl)
	fmt.Println(sl)

}