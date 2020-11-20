// dfs也称回溯算法, 常用于排列问题、组合问题、集合问题, 是一种时间复杂度很大的算法, 即也称为暴力算法
/*
dfs的算法框架为:
-----------------------------------
stack: 选择列表(一般为一个栈结构)
list: 可选列表(一般是原问题给定的数据)
cur: 当前处理的位置
res: 最终结果
-----------------------------------
if [满足条件]
	// 添加满足的结果到res中
	res.add(stack)
	return
for cur : list   // 遍历所有的位置
	// 使用当前选择(可能需要处理list的指针)
	stack.add(cur)
	// 递归调用, 计算选择cur时, 下一个结点的结果
	dfs(cur+1)
	// ****(重点) 回退之前的选择(也就是可以反悔)
	stack.pop()
*/

package dfs

import "example.com/algorithms/data_struct"

// 全排列问题
// 可以简化为cur位置可以选择list剩下的值,所以递归方程为:
/*
nums[i], nums[cur] = nums[cur], nums[i]
dfs(cur+1)
nums[i], nums[cur] = nums[cur], nums[i]
*/
func permute(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return [][]int{}
	}
	res := make([][]int, 0, l)
	var dfs func(cur int)
	dfs = func(cur int) {
		//注意: 递归终止条件是: 已选择到最后一个位置
		if cur >= l-1 {
			x := make([]int, l)
			copy(x, nums)
			res = append(res, x)
			return
		}
		for i := cur; i < l; i++ { //注意: i表示选择nums的i位置的数值,则可以选择l-1
			nums[i], nums[cur] = nums[cur], nums[i]
			dfs(cur + 1)
			nums[i], nums[cur] = nums[cur], nums[i]
		}
	}
	dfs(0)
	return res
}

/*
子集问题:
	子集问题和全排列的不同在于:
		1. 全排列是可以看做在一个与原nums等长的结果数组中的每个位置实验原所有数据;
		2. 子集问题是可以看作在原数组nums中的每个元素判断是否选择
		3. 在 "for cur : list "循环的表现就是:
			(1)全排列是for循环所有的剩余的list, 就是每个都必须取一次;
			(2)子集就是当前位置(cur)加入选择队列和不加入是两个情况, 分别递归;
*/
func subsets(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return [][]int{}
	}
	result := make([][]int, 0, 0)
	set := data_struct.NewSliceStack(0)
	var dfs func(int)
	dfs = func(cur int) {
		if cur >= l {
			// 需要把此时的栈状态保存,作为一个结果元素
			tmp := make([]int, len(set.Items()))
			copy(tmp, set.Items())
			result = append(result, tmp)
			return
		}
		// 把cur的值放入set中时, 此时cur+1用递归获得
		set.Push(nums[cur])
		dfs(cur + 1)

		// 把cur的值不放入set中时, 此时cur+1用递归获得
		set.Pop()
		dfs(cur + 1)
	}
	dfs(0)
	return result
}

/*
组合问题:  给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
	问题分析: 可以类比全排列问题, 只不过结果数组中只有两个元素, 在每个元素中尝试填入任意n组成的数组的元素
*/
func combine(n int, k int) [][]int {
	if n < 1 || k < 1 || n < k {
		return nil
	}
	result := make([][]int, 0, 0)
	choose := make([]int, 0, k)

	var dfs func(int)
	dfs = func(cur int) {
		if len(choose) == k || cur > n {
			if len(choose) != k {
				return
			}
			status := make([]int, k, k)
			copy(status, choose)
			result = append(result, status)
			return
		}
		// 选择cur
		choose = append(choose, cur)
		dfs(cur + 1)

		// 不选择cur
		choose = choose[:len(choose)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return result
}

// 目标和
func findTargetSumWays(nums []int, S int) int {
	left, res := S, 0
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			if left == 0 {
				res++
			}
			return
		}
		left = left - nums[cur]
		dfs(cur + 1)
		left = left + nums[cur]

		left = left + nums[cur]
		dfs(cur + 1)
		left = left - nums[cur]
	}
	dfs(0)
	return res
}
