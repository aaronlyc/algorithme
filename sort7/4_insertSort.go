package sort7

import "fmt"

//插入排序（排序10000个整数，用时约30ms）
func insertSort(data []int) {
	a, b := 0, len(data)
	insertionSort(data, a, b)
}

func insertionSort(data []int, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(data, j-1, j); j-- {
			swap(data, j-1, j)
			fmt.Printf("swap index %d and %d\n", j, j-1)
		}
	}
}
