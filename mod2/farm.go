package main

import (
	"fmt"
	"strings"
)

type IAnimal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{ name string }

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{ name string }

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hiss")
}

func create( farm map[string]IAnimal ) {
	var name, animal string

	var isAvailable = map[string]bool{
        "cow":   true,
        "bird":  true,
        "snake": true,
    }

	fmt.Println("\nEnter the name of the new animal\n>")
	fmt.Scan(&name)

	fmt.Println("\nEnter the type of the new animal (cow, bird or snake)\n>")
	fmt.Scan(&animal)
	animal = strings.ToLower(animal)

	if !isAvailable[animal]{
		fmt.Println("Type of animal unavaillable")
		return
	}

	switch animal {
	case "cow":
		farm[name] = &Cow{}
	case "bird":
		farm[name] = &Bird{}
	case "snake":
		farm[name] = &Snake{}
	}

	fmt.Println("Created it!")
}

func query( farm map[string]IAnimal ) {
	var name, info string
	fmt.Println("\nEnter the name of the animal\n>")
	fmt.Scan(&name)

	fmt.Println("\nEnter the info of the animal (Eat, Move, Speak)\n>")
	fmt.Scan(&info)
	info = strings.ToLower(info)
    
	animal, ok := farm[name]
    if !ok {
	   fmt.Println("Animal not found")
	   return
    }
	
	switch info {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid info. Try 'Eat', 'Move', or 'Speak'")
	}
}

func main() {

	farm := make(map[string]IAnimal, 3)

	for {
		var input string
		fmt.Println("Enter 'newanimal' to create or 'query' to get info\n>")
		fmt.Scan(&input)
		input = strings.ToLower(input)

		switch input {
		case "newanimal":
			create(farm)
		case "query":
			query(farm)
		default:
			fmt.Println("Try again")
		}
	}
}