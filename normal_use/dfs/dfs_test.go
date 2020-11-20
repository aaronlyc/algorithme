package dfs

import (
	"reflect"
	"testing"
	"time"
)

func Test_permute(t *testing.T) {
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
				nums: []int{1, 2, 3},
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				nums: []int{},
			},
			want: nil,
		},
		{
			name: "test3",
			args: args{
				nums: []int{1},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subsets(t *testing.T) {
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
				nums: []int{},
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				n: 1,
				k: 3,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				n: 3,
				k: 2,
			},
			want: [][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			name: "test2",
			args: args{
				n: 0,
				k: -2,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				n: 1,
				k: -2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(5 * time.Second)
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		n []int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				n: []int{1, 1, 1, 1, 1},
				k: 3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(5 * time.Second)
			if got := findTargetSumWays(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
