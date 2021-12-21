package main

import "fmt"

func main() {

	/* arrays */
	var arr1 [4]int
	var arr2 = [...]int{1, 5, 3, 6}

	fmt.Println(arr1)

	for index := range arr1 {
		arr1[index] = index
	}

	fmt.Println("new array ", arr1)
	fmt.Println(arr2)

	var arr3 = [...]int{1, 5, 3, 6}
	var arr4 [4]int = arr3

	arr3[0] = -1
	fmt.Println("arr3 ", arr3)
	fmt.Println("arr4 ", arr4)

	/* slice */

	var slc1 []int = []int{1, 2, 3, 4}
	var slc2 []int = make([]int, 4, 8)

	fmt.Println("slice1 ", slc1)
	fmt.Println("slice2 ", slc2)
	fmt.Println("slice1 ", slc1[:2])
	fmt.Println("Extended slice2 ", slc2[:cap(slc2)])
}
