package main

func fillArray(arr [10]int) [10]int {
	for i := range arr {
		arr[i] = i + 1
	}
	return arr
}
