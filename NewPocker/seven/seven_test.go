package seven

import "testing"

func Benchmark_SevenPocker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PokerMan()
	}
}
