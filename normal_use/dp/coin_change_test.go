package dp

import "testing"

func TestCoinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange1(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange2(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange3(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCoinChage(b *testing.B) {
	coins := []int{1, 2, 3, 5}
	amount := 30
	b.Run("coinChange1 bench", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			coinChange1(coins, amount)
		}
	})
	b.Run("coinChange2 bench", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			coinChange2(coins, amount)
		}
	})
	b.Run("coinChange3 bench", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			coinChange3(coins, amount)
		}
	})
}
