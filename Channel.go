package main

import "fmt"

func main() {
	var ch chan int = make(chan int, 1)
	multiplyByTwo(ch, 20)

	fmt.Println(<-ch)

	multiplyByTwo(ch, 10)
	fmt.Println(<-ch)

	multiplyByTwo(ch, 0)
	fmt.Println(<-ch)

}

func multiplyByTwo(ch chan int, value int) {
	ch <- value * 2
}
