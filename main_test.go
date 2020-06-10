package main

import (
	"testing"
)

func BenchmarkSolvePuzzle(b *testing.B) {
	puzzle = [9][9]int{
		{0, 0, 0, 0, 0, 6, 0, 9, 0},
		{0, 5, 0, 3, 4, 9, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 5},
		{0, 0, 0, 7, 0, 0, 0, 0, 8},
		{0, 3, 0, 0, 0, 8, 0, 6, 0},
		{0, 0, 5, 1, 3, 0, 0, 7, 0},
		{1, 9, 7, 0, 0, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 5, 6, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	for i := 0; i < b.N; i++ {
		fillInSingleCandidate()
		solve()
	}
}
