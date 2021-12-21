package main

import "fmt"

func main() {

	var i int
	for i = 0; i < 5; i++ {
		fmt.Println("looping up for ", i, " times ")
	}

	for i > 0 {
		fmt.Println("looping down to ", i)
		i--
	}
}
