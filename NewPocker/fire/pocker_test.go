package fire

import (
	"testing"
)

func Benchmark_Pocker5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PokerMan()
	}
}