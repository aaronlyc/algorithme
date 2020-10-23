package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	if nums[len(nums)-1] <= 0 {
		return nil
	}
	res := make([][]int, 0, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else if sum > 0 {
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			}
		}
	}
	return res
}

/*
注意点:
(1) 首先需要排序(从小到大排序, 时间复杂度: O(nlogn), 空间复杂度: O(1))
(2) 边界条件:
	1> len(nums)>=3;
	2> min(nums)<0;
	3> max(nums)>0;
(3) 循坏nums获取第一个数字下标i, 则第二个数为i+1, 第三个数为len(nums)-1
(4) 去重:
	1> 第一个数去重: i>0 && nums[i]==nums[i-1], continue
	2> 第二个数去重: l<r && nums[l]==nums[l+1], l++; l++
	3> 第三个数去重: l<r && nums[r]==nums[r-1], r--; r--
(5) 把满足要求的结果保存在二维数组中
*/
