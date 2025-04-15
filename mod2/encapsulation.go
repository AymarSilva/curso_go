package main

import "fmt"

type Animal struct{
	food string
	locomotion string
	noise string
}

func (req *Animal) Eat() string {
	return req.food
}

func (req *Animal) Move() string {
	return req.locomotion
}

func (req *Animal) Speak() string {
	return req.noise
}

func main(){

	animals := map[string]Animal{
		"cow" : {"grass","walk","moo"},
		"bird": {"worms","fly","peep"},
		"snake": {"mice","slither","hsss"},
	}

	actions := map[string]func(*Animal) string{
		"eat":   (*Animal).Eat,
		"move":  (*Animal).Move,
		"speak": (*Animal).Speak,
	}

	var name_animal, action_animal string

	for{
		fmt.Println("animals: cow, bird, snake")
		fmt.Println("action: eat, move, speak")

		fmt.Print("Write animal name and its action: ")
		fmt.Scan(&name_animal, &action_animal)

		animal:= animals[name_animal]
		action:= actions[action_animal]

		fmt.Println( action(&animal) )
	}
}