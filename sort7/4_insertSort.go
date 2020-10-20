package sort7

//插入排序（排序10000个整数，用时约30ms）
func insertionSort(data []int, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(data, j-1, j); j-- {
			swap(data, j-1, j)
			//fmt.Printf("swap index %d and %d\n", j, j-1)
		}
	}
}
