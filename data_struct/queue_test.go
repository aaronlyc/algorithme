package data_struct

import "testing"

func Test_queue(t *testing.T) {
	q := NewSliceQueue()
	if q.Empty() {
		t.Log("q is empty")
	}
	q.Poll()

	q.Offer(23)
	q.Poll()

	if q.Empty() {
		t.Log("q is empty")
	}
}
