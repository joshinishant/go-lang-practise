package main

import (
	"errors"
	"fmt"
)

func main() {

	var marks int = 70
	grade, err := getGrade(marks)

	if err != nil {
		panic(err)
	}

	printGrade(grade)

}

func getGrade(marks int) (string, error) {

	var grade string

	if marks < 0 {
		return "g", errors.New("Marks obtained cannot be negative")
	}

	if marks >= 90 {
		grade = "A"
	} else if marks >= 75 {
		grade = "B"
	} else if marks >= 60 {
		grade = "C"
	} else if marks >= 35 {
		grade = "D"
	} else {
		grade = "F"
	}

	return grade, nil
}

func printGrade(grade string) {
	switch grade {
	case "A":
		fmt.Println("Passed with Distinction")
	case "B":
		fmt.Println("Passed with First Class")
	case "C":
		fmt.Println("Passed with Second Class")
	case "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Grade not known")
	}
}
