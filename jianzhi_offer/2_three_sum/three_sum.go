package __three_sum

import "sort"

func threeSum(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		tRes := twoSum(nums, i+1, -1*cur)
		if len(tRes) > 0 {
			for j := 0; j < len(tRes); j++ {
				tRes[j] = append(tRes[j], cur)
				result = append(result, tRes[j])
			}
		} else {
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return result
}

// 从start开始的[]int中找到目标值为target的两数集合
func twoSum(sortedNums []int, start, target int) [][]int {
	var result [][]int
	end := len(sortedNums) - 1
	for start < end {
		sum := sortedNums[start] + sortedNums[end]
		left, right := sortedNums[start], sortedNums[end]
		if sum < target {
			for start < end && sortedNums[start] == left {
				start++
			}
		} else if sum > target {
			for start < end && sortedNums[end] == right {
				end--
			}
		} else {
			result = append(result, []int{sortedNums[start], sortedNums[end]})
			for start < end && sortedNums[start] == left {
				start++
			}
			for start < end && sortedNums[end] == right {
				end--
			}
		}
	}
	return result
}
