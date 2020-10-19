package sort7

import (
	"math/rand"
	"testing"
	"time"
)

func Test_heapSort(t *testing.T) {
	type args struct {
		data []int
		a    int
		b    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				data: []int{10, 7, 6, 8, 9, 5, 4},
				a:    0,
				b:    7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("the original nums is: %v\n", tt.args.data)
			heapSort(tt.args.data, tt.args.a, tt.args.b)
			t.Logf("the sorted nums is: %v\n", tt.args.data)
		})
	}
}

func Benchmark_heapSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var testSlice [10000]int
	for i := 0; i < 10000; i++ {
		testSlice[i] = rand.Int()
	}
	var test []int
	test = testSlice[:]
	b.ResetTimer()

	b.Run("test1", func(b *testing.B) {
		heapSort(test, 0, 10000)
	})

	b.Log("ok success")
}
