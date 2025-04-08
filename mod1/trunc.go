package main

import ("fmt")

var number float32

func truncateNumber(){
	fmt.Print("Write it down: ")
	fmt.Scan(&number)
	newNumber:= int32(number)
	fmt.Println(newNumber)
}

func main(){
	truncateNumber()
}