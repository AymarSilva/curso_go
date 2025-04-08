package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Address string
}

func main()  {
	var name, address string
	mapy:= make(map[string]string)


	fmt.Printf("Write your name: ")
	fmt.Scan(&name)
	mapy["Name"] = name

	fmt.Printf("Write your address: ")
	fmt.Scan(&address)
	mapy["Address"] = address

	person:= Person{
		Name: mapy["Name"],
		Address: mapy["Address"],
	}

	data, err:= json.Marshal(person)

	if err != nil {
		fmt.Println("JSON error: ",err)
		return
	}

	fmt.Println("JSON: ", string(data))
}