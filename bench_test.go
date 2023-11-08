package sha1

import (
	"crypto/sha1"
	"testing"
)

// BenchmarkCustomSha1_2000 measures the performance of the Sum function for a 2000-byte input.
func BenchmarkCustomSha1_2000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randInput := randSeq(2000)
		Sum([]byte(randInput))
	}
}

// BenchmarkBuiltinSha1_2000 measures the performance of the built-in crypto/sha1 package for a 2000-byte input.
func BenchmarkBuiltinSha1_2000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randInput := randSeq(2000)
		sha1.Sum([]byte(randInput))
	}
}

// BenchmarkCustomSha1_20000 measures the performance of the Sum function for a 20000-byte input.
func BenchmarkCustomSha1_20000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randInput := randSeq(20000)
		Sum([]byte(randInput))
	}
}

// BenchmarkBuiltinSha1_20000 measures the performance of the built-in crypto/sha1 package for a 20000-byte input.
func BenchmarkBuiltinSha1_20000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randInput := randSeq(20000)
		sha1.Sum([]byte(randInput))
	}
}
