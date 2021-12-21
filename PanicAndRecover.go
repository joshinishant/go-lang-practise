package main

import "fmt"

/*func main() {
	AllOperations(10, 0)
	fmt.Println("Exiting main method")
}*/

func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println("Recovering from panic due to - ", err)
	} else {
		fmt.Println("Execution completed")
	}
}

func AllOperations(p1, p2 int) {
	defer handlePanic()
	fmt.Println("Parameter 1 - ", p1, ", Parameter 2 - ", p2)
	fmt.Println("Addition - ", p1+p2)
	fmt.Println("Substraction - ", p1-p2)
	fmt.Println("Multiplication - ", p1*p2)
	fmt.Println("Division - ", p1/p2)

}
