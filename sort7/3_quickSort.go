package sort7

//快速排序（排序10000个随机整数，用时约0.9ms）
func quickSort(nums []int) {
	recursionSort(nums, 0, len(nums)-1)
}

func recursionSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	recursionSort(nums, left, pivot-1)
	recursionSort(nums, pivot+1, right)
}

/*
partition 为快速排序的切分算法, 每切分一次可以确定切分点的位置, 这里主要是二向切分, 下面还会有三向切分
*/

func partition(nums []int, left int, right int) int {
	i, j := left, right
	for {
		for i < j && nums[left] <= nums[j] {
			j--
		}
		for i < j && nums[left] >= nums[i] {
			i++
		}
		if i >= j {
			break
		}
		swap(nums, i, j)
	}
	swap(nums, left, j)
	return i
}

//func partition(nums []int, left int, right int) int {
//	for left < right {
//		for left < right && nums[left] <= nums[right] {
//			right--
//		}
//		if left < right {
//			nums[left], nums[right] = nums[right], nums[left]
//			left++
//		}
//
//		for left < right && nums[left] <= nums[right] {
//			left++
//		}
//		if left < right {
//			nums[left], nums[right] = nums[right], nums[left]
//			right--
//		}
//	}
//
//	return left
//}
