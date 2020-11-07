package dp

import (
	"testing"
)

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{n: 10},
			want: 55,
		},
		{
			name: "test2",
			args: args{n: -10},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib1(tt.args.n); got != tt.want {
				t.Errorf("Fib1ToShowNode() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := fib2(tt.args.n); got != tt.want {
				t.Errorf("Fib2ToShowNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
BenchmarkFib1/Fib1-8               44444             27995 ns/op               0 B/op          0 allocs/op
BenchmarkFib1/Fib2-8            13182553                84.7 ns/op           176 B/op          1 allocs/op
BenchmarkFib1/Fib3-8            100000000               10.8 ns/op             0 B/op          0 allocs/op
*/
func BenchmarkFib(b *testing.B) {
	b.Run("Fib1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fib1(20)
		}
	})
	b.Run("Fib2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fib2(20)
		}
	})
	b.Run("Fib3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fib3(20)
		}
	})
}
