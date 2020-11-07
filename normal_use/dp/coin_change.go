/*
当 coins = []int{1,2,3,5}  amount = 30 时:
BenchmarkCoinChage/coinChange1_bench-8                 1        2093032400 ns/op               0 B/op          0 allocs/op
BenchmarkCoinChage/coinChange2_bench-8            200002              5895 ns/op            2232 B/op          6 allocs/op
BenchmarkCoinChage/coinChange3_bench-8           1938613               614 ns/op             256 B/op          1 allocs/op
综上可知: 使用 dp table 的方式会得到最高性能
*/
package dp

import "math"

/**
 * @Author liuyuchao
 * @Date 13:44 2020/11/7
 * @Param coins 不同面值的硬币; amount 目标总金额
 * @return int 可以凑成总金额所需的最少的硬币个数
 * @Description "给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
 *              如果没有任何一种硬币组合能组成总金额，返回-1. coinChange1 是包含重复子问题的解法"
 **/
func coinChange1(coins []int, amount int) int {
	// 状态转移方程的思路:
	// 1. 确定状态的含义: 状态是 amount 值;
	// 2. 确定选择: 在 coins 中选择一个能得到 amount 值;
	// 3. 明确 base case: 当 amount<0 返回-1, 当 amount=0, 返回0;
	// 4. 定义dp数组/函数: 输入 amount, 返回能凑成当前 amount 时最少的硬币数量.
	var dp func(int) int
	dp = func(n int) int {
		if n < 0 {
			return -1
		}
		if n == 0 {
			return 0
		}
		res := math.MaxInt32
		for i := 0; i < len(coins); i++ {
			sub := dp(n - coins[i])
			if sub == -1 {
				continue
			}
			res = min(res, sub+1)
		}
		if res == math.MaxInt32 {
			return -1
		}
		return res
	}
	return dp(amount)
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// coinChange2 使用备忘录方式优化, 去除重复子问题
func coinChange2(coins []int, amount int) int {
	// 状态转移方程的思路:
	// 1. 确定状态的含义: 状态是 amount 值;
	// 2. 确定选择: 在 coins 中选择一个能得到 amount 值;
	// 3. 明确 base case: 当 amount<0 返回-1, 当 amount=0, 返回0;
	// 4. 定义dp数组/函数: 输入 amount, 返回能凑成当前 amount 时最少的硬币数量.
	var dp func(int) int
	memo := make(map[int]int) // memo 是备忘录的方式记录不同的n时, 最少的硬币数量
	dp = func(n int) int {
		// 如果备忘录中已有值, 直接返回, 防止重复计算
		if v, ok := memo[n]; ok {
			return v
		}
		if n < 0 {
			return -1
		}
		if n == 0 {
			return 0
		}
		res := math.MaxInt32
		for i := 0; i < len(coins); i++ {
			sub := dp(n - coins[i])
			if sub == -1 {
				continue
			}
			res = min(res, sub+1)
		}
		if res == math.MaxInt32 {
			return -1
		}
		memo[n] = res
		return res
	}
	return dp(amount)
}

// coinChange3 使用dp table 方式优化, 去除重复子问题, 并且自底向上
func coinChange3(coins []int, amount int) int {
	// 状态转移方程的思路:
	// 1. 确定状态的含义: 状态是 amount 值;
	// 2. 确定选择: 在 coins 中选择一个能得到 amount 值;
	// 3. 明确 base case: 当 amount<0 返回-1, 当 amount=0, 返回0;
	// 4. 定义dp数组/函数: 输入 amount, 返回能凑成当前 amount 时最少的硬币数量.
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	dpTable := make([]int, amount+1, amount+1)
	dpTable[0] = 0
	for i := 1; i < len(dpTable); i++ {
		for coin := 0; coin < len(coins); coin++ {
			if i-coins[coin] < 0 {
				continue
			}
			if i-coins[coin] != 0 && dpTable[i-coins[coin]] == 0 {
				continue
			}
			dpTable[i] = minDp(dpTable[i], 1+dpTable[i-coins[coin]])
		}
	}
	if dpTable[amount] == 0 {
		return -1
	}
	return dpTable[amount]
}

func minDp(a, b int) int {
	switch {
	case a == 0:
		return b
	case b == 0:
		return a
	case a <= b:
		return a
	default:
		return b
	}
}
