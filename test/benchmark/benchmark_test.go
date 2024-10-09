package benchmark

import "testing"

func BenchmarkCreateStringWithoutStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateStringWithoutStringBuilder()
	}
}

func BenchmarkCreateStringWithStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateStringWithStringBuilder()
	}
}
//$ go test -bench=.
//$ go test -bench=. -benchmem

