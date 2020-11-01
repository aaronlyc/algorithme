package data_struct

// 可以用数组简单实现
type queue struct {
	item []interface{}
}

func NewSliceQueue() *queue {
	return &queue{item: make([]interface{}, 0, 0)}
}

func (q *queue) Offer(n interface{}) {
	q.item = append(q.item, n)
}

func (q *queue) Empty() bool {
	return len(q.item) == 0
}

func (q *queue) Poll() interface{} {
	if q.Empty() {
		return nil
	}
	res := q.item[0]
	q.item = q.item[1:]
	return res
}

func (q *queue) Size() int {
	return len(q.item)
}
