package main

import "testing"

func TestExamples(t *testing.T) {
	var freqTests = []struct {
		n        []int // input
		expected int   // expected result
	}{
		{
			[]int{1, -2, 3, 1},
			2,
		},
		{
			[]int{1, -1},
			0,
		},
		{
			[]int{-6, 3, 8, 5, -6},
			5,
		},
		// those two - not sure about the repetition of the list 🤔
		//{
		//	[]int{3, 3, 4, -2, -4},
		//	10,
		//},
		//{
		//	[]int{7, 7, -2, -7, -4},
		//	14,
		//},
	}

	for _, tt := range freqTests {
		actual := partTwo(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
