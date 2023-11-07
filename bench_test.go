package sha1

import (
	"crypto/sha1"
	"testing"
)

func BenchmarkCustomSha1_2000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randInput := randSeq(2000)
		Sum([]byte(randInput))
	}
}

func BenchmarkBuiltinSha1_2000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randInput := randSeq(2000)
		sha1.Sum([]byte(randInput))
	}
}

func BenchmarkCustomSha1_20000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randInput := randSeq(20000)
		Sum([]byte(randInput))
	}
}

func BenchmarkBuiltinSha1_20000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randInput := randSeq(20000)
		sha1.Sum([]byte(randInput))
	}
}
