package main

import "testing"

func Test_poker5(t *testing.T) {
	PokerMan()
}

func Benchmark_poker5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PokerMan()
	}
}