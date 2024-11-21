package main

import "testing"

// BenchmarkFib1-16		413528		2934 ns/op		0 B/op		0 allocs/op
func BenchmarkFib1(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Fib1(15)
	}
}

// BenchmarkFib2-16		551952531		2.152 ns/op		0 B/op		0 allocs/op
func BenchmarkFib2(b *testing.B) {
	n := 15
	memo := make([]int, n)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Fib2(n, memo)
	}
}

// BenchmarkFib3-16        100000000               10.90 ns/op            0 B/op		0 allocs/op
func BenchmarkFib3(b *testing.B) {
	n := 15
	k := make([]int, n+1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Fib3(n, k)
	}
}
