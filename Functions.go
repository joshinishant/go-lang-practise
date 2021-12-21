package main

import (
	"errors"
	"fmt"
)

func main() {

	result, _ := f2("Test")

	var err error

	if result, err = f2(""); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result = f1(10, 20, 65)
	fmt.Println(result)

	func() {
		fmt.Println("Anonymous function without parameter ")
	}()

	func(p1 string) {
		fmt.Println("Anonymous function with parameter ", p1)
	}("Hello")

}

func f1(p1, p2 int, p3 float64) string {
	fmt.Println("Parameter 1 ", p1, " Parameter 2 ", p2, " Parameter 3 ", p3)
	return "complete"
}

func f2(p1 string) (string, error) {
	if len(p1) <= 0 {
		return "", errors.New("String cannot be empty")
	}

	return p1, nil
}
