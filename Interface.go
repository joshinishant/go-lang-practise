package main

import (
	"errors"
	"fmt"
)

type Animal interface {
	speak(count int) (string, error)
}

type Dog struct{}

func (dog Dog) speak(count int) (string, error) {
	if count <= 0 {
		return "", errors.New("count shoulb be a positive number")
	} else if count == 1 {
		return "bow", nil
	} else {
		return "boooooooow", nil
	}
}

type Cat struct{}

func (cat Cat) speak(count int) (string, error) {
	if count <= 0 {
		return "", errors.New("count shoulb be a positive number")
	} else if count == 1 {
		return "meow", nil
	} else {
		return "meooooooooow", nil
	}
}

func sound(animals Animal, count int) {
	if result, err := animals.speak(count); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func main() {
	var animalList []Animal = []Animal{Cat{}, Dog{}}

	for _, value := range animalList {
		sound(value, 3)
	}
}
