package sevenJoker

import "testing"

func Benchmark_pokerMan(b *testing.B) {
	for i:=0; i< b.N; i++ {
		PokerMan()
	}
}