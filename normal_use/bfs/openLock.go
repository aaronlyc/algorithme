package bfs

import (
	"context"
	"example.com/algorithms/data_struct"
	"sync"
)

// 可以开两个goroutine, 分别从起点"0000"和终端target进行双向bfs
// 共享变量: visited, 存两个goroutune访问过的,需要RWMutex
func openLock(deadends []string, target string) int {
	var visited sync.Map
	stepChan := make(chan int)
	for _, deadend := range deadends {
		visited.Store(deadend, struct{}{})
	}
	var bfs func(context.Context, string, string)
	ctx, cancelFunc := context.WithCancel(context.TODO())
	bfs = func(ctx context.Context, start, target string) {
		select {
		case <-ctx.Done():
			return
		default:
			q := data_struct.NewSliceQueue()
			q.Offer(start)
			visited.Store(start, struct{}{})
			var step int
			for !q.Empty() {
				sz := q.Size()
				for i := 0; i < sz; i++ {
					cur := q.Poll()
					if cur == target {
						stepChan <- step
						return
					}
					for j := 0; j < 4; j++ {
						visited.LoadOrStore(plusOne(cur.(string), j), struct{}{})
						visited.LoadOrStore(minusOne(cur.(string), j), struct{}{})
					}
				}
				step++
			}
			stepChan <- -1
			return
		}
	}
	go bfs(ctx, "0000", target)
	go bfs(ctx, target, "0000")
	select {
	case step := <-stepChan:
		cancelFunc()
		return step
	}
}

func plusOne(from string, i int) string {
	cur := []byte(from)
	c := cur[i] + '0'
	if c == '9' {
		c = '0'
	} else {
		c = c + 1
	}
	return string(cur)
}

func minusOne(from string, i int) string {
	cur := []byte(from)
	c := cur[i] + '0'
	if c == '0' {
		c = '9'
	} else {
		c = c - 1
	}
	return string(cur)
}
