package sort7

func shellSort(nums []int) {
	a, b := 0, len(nums)
	insertionSort(nums, a, b)
}
