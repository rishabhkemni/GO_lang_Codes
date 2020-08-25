package main

import (
	"fmt"
	_ "strings"
)

type Animal struct {
	food string
	locomotion string
	sound string
}

func (animal Animal) Eat() {
	fmt.Println("eats :-", animal.food)
}

func (animal Animal) Move() {
	fmt.Println("moves :-", animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println("sounds like :-", animal.sound)
}

func main() {
	var animalMapping map[string]Animal
	animalMapping = make(map[string]Animal)
	animalMapping["cow"] = Animal{"grass", "walk", "moo"}
	animalMapping["bird"] = Animal{"worms", "fly", "peep"}
	animalMapping["snake"] = Animal{"mice", "slither", "hisss"}

	for true {
		var animalName, action string
		fmt.Print("--RK-->")
		fmt.Scanln(&animalName, &action)
		switch action {
			case "eat" :
				animalMapping[animalName].Eat()
				break
			case "move" :
				animalMapping[animalName].Move()
				break
			case "speak" :
				animalMapping[animalName].Speak()
        break
			default :
				fmt.Println("Wrong Arguments Please Try Again !!!")
        fmt.Println("Arguments should be <animal name> <eat|speak|move>")
		}
	}
}
