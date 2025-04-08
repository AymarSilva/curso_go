package main

import (
	"fmt"
	"strings"
)

func searchars(text string) bool {
	isThereA := strings.Contains(text, "a")
	hasPrefixI := strings.HasPrefix(text, "i")
	hasSuffixN := strings.HasSuffix(text, "n")

	if isThereA && hasPrefixI && hasSuffixN{
		return true
	}else{
		return false
	}
}

func main()  {
	var text string

	fmt.Print("Enter the text: ")
	fmt.Scan(&text)
	text = strings.ToLower(text)
	
	if searchars(text){
		fmt.Println("Found")
	}else{
		fmt.Println("Not Found")
	}
}