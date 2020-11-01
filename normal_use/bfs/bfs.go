/*
BFS: 一般用来寻找最短路径问题, 需要用到queue数据结构, 走迷宫, 解密码锁等会用到
BFS算法框架为:
func bfs(start, target)
	var q queue //核心数据结构, 放入需要处理的结点
	visited := make(map[*node]struct{}) //存入已遍历的结点, 防止走回头路

	//首先加入start结点
	q.offer(start)
	visited.add(start)

	var step int //存入走过的步数, 本实例就是求步数,其他问题需要结合题目

	//遍历q队列的所有结点
	for (q not empty)
		sz := q.size //获取q的长度(注意: 需要现在获取, 不然后面q的长度会变化,需要得到现在的状态)
		//拿出q队列中的所有结点
		for	(i:[0,sz))
			cur := q.poll()  //取出除当前的处理
			//第一步: 先判断当前结点是否满足target要求
			if (cur is target)
				return step   //这里假设需要返回step,其他问题需要结合题意考虑
			//第二步: 把cur相邻的其他结点放入q队列中
			for (x:cur.adj())  //cur.adj()表示cur的周围相邻结点
				if (x not in visited) //判断x是否已被访问过, 如果已经访问过了, 则不需要加入了
					q.offer(x)
					visited.add(x)
		//第三步: 步数加1(该层都遍历完)
		step++
*/
package bfs

import "example.com/algorithms/data_struct"

// bfs 此处为bfs的一般模板. 主要思想是: 用q保存需要访问的队列, visited记录已访问过的结点
func bfs(start, target int) int {
	q := data_struct.NewSliceQueue()
	visited := make(map[int]struct{})

	q.Offer(start)
	visited[start] = struct{}{}

	var step int

	for !q.Empty() {
		sz := q.Size()

		for i := 0; i < sz; i++ {
			cur := q.Poll()
			if cur == target {
				return step
			}

			//cur周围的结点
			adjs := adj(cur)
			for j := 0; j < len(adjs); j++ {
				// 判断是否已经访问过了
				if _, ok := visited[adjs[j].(int)]; !ok {
					q.Offer(adjs[j])
					visited[adjs[j].(int)] = struct{}{}
				}
			}
		}
		step++
	}
	return step
}

// cur周围结点组成的数组
func adj(cur interface{}) []interface{} {
	//todo: 根据实际问题写逻辑
	return []interface{}{}
}
