package sha1

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	tests := []string{
		"",
		"Hello World!",
		"Distributed Lab",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}
	for i := 0; i < 10000; i++ {
		length := rand.Intn(20000)
		tests = append(tests, randSeq(length))
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Hash %s\n", tt), func(t *testing.T) {
			want := sha1.Sum([]byte(tt))
			if got := Sum([]byte(tt)); !reflect.DeepEqual(got, want[:]) {
				t.Errorf("Hash() = %v, want %v", got, want)
			}
		})
	}
}
