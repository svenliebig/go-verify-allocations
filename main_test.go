package main

import "testing"

func BenchmarkAllocations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
	b.ReportAllocs()
}
