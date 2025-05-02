package main

import (
	"fmt"
	sort "golang-learning/algorithm/sort"
	helper "golang-learning/helper"
)

func main() {
	var slice []int = helper.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sort.Insertionsort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
