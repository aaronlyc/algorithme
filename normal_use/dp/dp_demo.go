package dp

// fib 是典型的斐波拉契数列

// fib1 是原始的递归解法
// 递归方程式为:
//	1, n=1,2;
//  f(n-1)+f(n-2), n>2;
func fib1(n int) int {
	if n < 1 {
		return -1
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// fib2 是使用dp数组方式解决重复子问题
func fib2(n int) int {
	if n < 1 {
		return -1
	}
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n+1, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

// fib3 是使用状态压缩方法, 使用两个辅助变量记录前面两个数值
/**
 * @Author liuyuchao
 * @Date 1:15 2020/11/7
 * @Param
 * @return
 * @Description "描述一个问题, 来表示一个不错的问题, 这个问题需要很多方的参与"
"
 **/
func fib3(n int) int {
	if n < 1 {
		return -1
	}
	if n == 1 || n == 2 {
		return 1
	}
	first, second, res := 1, 1, 0
	for i := 3; i <= n; i++ {
		res = first + second
		first = second
		second = res
	}
	return res
}
