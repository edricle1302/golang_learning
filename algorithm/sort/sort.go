package sort

func Insertionsort(items []int) []int {
	newSlice := items
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if newSlice[j-1] > newSlice[j] {
				newSlice[j-1], newSlice[j] = newSlice[j], newSlice[j-1]
			}
			j = j - 1
		}
	}
	return newSlice
}
