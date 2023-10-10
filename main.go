package main

import (
	"fmt"
)

func main() {
	var slice []int = helper.generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sort.insertionsort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
