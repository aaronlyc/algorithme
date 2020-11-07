package sort7

import "testing"

func Test_quickSort(t *testing.T) {
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
				[]int{5, 5, 3, 3, 2, 2, 8, 8, 0, 0, 1, 1},
			},
		},
		{
			name: "test2",
			args: args{
				[]int{1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.nums)
			quickSort(tt.args.nums)
			t.Log(tt.args.nums)
		})
	}
}
