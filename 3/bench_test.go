package main

import (
	"testing"
)


var test_slice = GenerateSlice(1000)


func BenchmarkQuick(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		QuickSort(test_slice)
	}
}

func BenchmarkInsert(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		InsertionSort(test_slice)
	}
}

func BenchmarkBubble(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		BubbleSort(test_slice)
	}
}
