package sort7

import (
	"testing"
)

func Test_shellSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				[]int{5, 4, 3, 8, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shellSort(tt.args.nums)
			t.Logf("sorted nums is:%v\n", tt.args.nums)
		})
	}
}
