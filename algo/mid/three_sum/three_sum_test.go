package three_sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				[]int{},
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				[]int{1},
			},
			want: nil,
		},
		{
			name: "test3",
			args: args{
				[]int{0, -1, -2},
			},
			want: nil,
		},
		{
			name: "test4",
			args: args{
				[]int{0, -1, -2},
			},
			want: nil,
		},
		{
			name: "test5",
			args: args{
				[]int{0, -1, -2, 1, 2},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
