package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input          string
	expected int
}{
	{"1212", 6},
	{"1221", 0},
	{"123425", 4},
	{"123123", 12},
	{"12131415", 4},
}

func TestGivenInputsAndOutputs(t *testing.T) {
	var result int

	for _, tc := range testCases {

		result = findSumConvertAsNeeded(tc.input)
		assert.Equal(t, tc.expected, result)

		result = findSumConvertEverything(tc.input)
		assert.Equal(t, tc.expected, result)

		result = findSumCleverConvertAsNeeded(tc.input)
		assert.Equal(t, tc.expected, result)

		result = findSumCleverConvertEverything(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}

var benchmarkCases = []struct {
	name  string
	input string
}{
	{"Short", "12131415"},
	{"Medium", "12133232141532177217565631415121121332321415321772175656314153323214153217712133232141532177217565631415217565631415"},
	{"Long", puzzleInput},
}

func BenchmarkFindSumCompareStringsConvertToIntAsNeeded(b *testing.B) {
	for _, bm := range benchmarkCases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findSumConvertAsNeeded(bm.input)
			}
		})
	}
}

func BenchmarkFindSumConvertEverythingToIntBeforeComparing(b *testing.B) {
	for _, bm := range benchmarkCases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findSumConvertEverything(bm.input)
			}
		})
	}
}

func BenchmarkFindSumConvertAsNeededStopHalfway(b *testing.B) {
	for _, bm := range benchmarkCases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findSumCleverConvertAsNeeded(bm.input)
			}
		})
	}
}

func BenchmarkFindSumConvertEverythingStopHalfway(b *testing.B) {
	for _, bm := range benchmarkCases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findSumCleverConvertEverything(bm.input)
			}
		})
	}
}