package sort7

func less(nums []int, i, j int) bool {
	return nums[i] < nums[j] // 表示需要从小到大排序
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
