package sort7

import (
	"math/rand"
	"testing"
	"time"
)

func Test_insertSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func BenchmarkInsertSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var testSlice [10000]int
	for i := 0; i < 10000; i++ {
		testSlice[i] = rand.Int()
	}
	var test []int
	test = testSlice[:]
	b.ResetTimer()

	b.Run("test1", func(b *testing.B) {
		insertionSort(test, 0, 10000)
	})
}
