package mid

import (
	"example.com/algorithms/mock"
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"test1", []int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := mock.MackIntSlice2List(tt.args)
			want := mock.MackIntSlice2List(tt.want)
			if got := sortList(input); !reflect.DeepEqual(got, want) {
				t.Errorf("SortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
