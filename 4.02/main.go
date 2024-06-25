package main

import "fmt"

func compArrays() (bool, bool, bool, bool) {
	var arr1 [5]int
	var arr2 [5]int
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	var arr5 [5]int
	return arr1 == arr2, arr1 == arr3, arr1 == arr4, arr1 == arr5
}

func main() {
	comp1, comp2, comp3, comp4 := compArrays()
	fmt.Println("arr1 == arr2:", comp1)
	fmt.Println("arr1 == arr3:", comp2)
	fmt.Println("arr1 == arr4:", comp3)
	fmt.Println("arr1 == arr5:", comp4)
}
