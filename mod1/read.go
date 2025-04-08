package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fName string
	lName string
}

func main() {
	var fName, lName string
	
	fmt.Println("Write first name: ")
	fmt.Scan(&fName)
	
	fmt.Println("Write last name: ")
	fmt.Scan(&lName)

	nameFile,_ := os.OpenFile("names.txt", os.O_APPEND|os.O_CREATE, 0644)
	defer nameFile.Close()

	nameFile.WriteString(fName + " " + lName + "\n")

	nameFile,_ = os.Open("names.txt")
	defer nameFile.Close()

	byteArray,_ := os.ReadFile("names.txt")

	parts := strings.Fields(string(byteArray))
	names := make([]Person, 0)

	for i := 0; i < len(parts); i += 2 {
		if i+1 < len(parts) {
			names = append(names, Person{
				fName: parts[i],
				lName: parts[i+1],
			})
		}
	}

	for _, person := range names {
		fmt.Println(person.fName, person.lName)
	}
}
