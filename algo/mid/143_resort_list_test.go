package mid

import (
	"example.com/algorithms/algo/mock"
	"reflect"
	"testing"
)

func Test_reorderList(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"test1", []int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := mock.MackIntSlice2List(tt.args)
			want := mock.MackIntSlice2List(tt.want)
			if reorderList(input); !reflect.DeepEqual(input, want) {
				var result []int
				var i int
				for input != nil {
					i++
					result = append(result, input.Val)
					if i > len(tt.want) {
						t.Error("循环链表了")
						return
					}
				}
				t.Errorf("SortList() = %v, want %v", result, tt.want)
			}
		})
	}
}
