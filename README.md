# SHA-1 (Secure Hash Algorithm 1)

SHA-1 implementation in Golang. Task №5 for the Cryptography for Developers course.

Ref: https://en.wikipedia.org/wiki/SHA-1

## How to use

To use this SHA-1 implementation, execute the following command first:
```bash
go get github.com/danielost/sha-1
```
Then simply add the following import to your Go code:
```golang
import (
	// Importing the SHA-1 package
	sha1 "github.com/danielost/sha-1"
)
```
Now it's possible to use the `Sum([]byte) []byte` function, which takes a byte array of an arbitrary length and returns a fixed-size message digest:
```golang
// The message for which we want to calculate the SHA-1 hash
msg := "Cats control the outer ear using 32 muscles; humans use 6"

// Calculate the SHA-1 hash of the message.
// Sum is represented as a byte array of length 20 (160 bits)
msgDigest := sha1.Sum([]byte(msg))
```
> The fact that the function takes a byte array as an input means it's possible to hash any data, from text to audio or video files.

Full block of code, that follows the above steps:
```golang
package main

import (
	"fmt"

	// Importing the SHA-1 package
	sha1 "github.com/danielost/sha-1"
)

func main() {
	// The message for which we want to calculate the SHA-1 hash
	msg := "Cats control the outer ear using 32 muscles; humans use 6"

	// Calculate the SHA-1 hash of the message.
	// Sum is represented as a byte array of length 20 (160 bits)
	msgDigest := sha1.Sum([]byte(msg))

	// Print the resulting hash in hexadecimal format
	fmt.Printf("%x\n", msgDigest)
}

```

## Running tests and benchmarks

**1. Tests:**

The tests are located in `./hash_test.go`. The same input is hashed using this SHA-1 implementation and the [built-in SHA-1](https://pkg.go.dev/crypto/sha1). Then their outputs are compared. The test data set contains some hardcoded values as well as 10000 randomly generated sequences.

To run the tests, clone the repository and execute the following command:
```bash
go test .
```
If the program works correctly, the output will look as follows:
```bash
ok      github.com/danielost/sha-1      4.558s
```

**2. Benchmarks**

Benchmarks were added to compare the execution time of two implementations.
To run the benchmark, clone the repository and execute the following command:
```bash
go test -bench=.
```
This command actually runs both the tests and benchmarks, so there is no need to run the tests separately if you are interested in running everything.

The output of the above command should look like this:
```bash
goos: linux
goarch: amd64
pkg: github.com/danielost/sha-1
cpu: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
BenchmarkCustomSha1_2000-2     	   14986	     77941 ns/op
BenchmarkBuiltinSha1_2000-2    	   19896	     59004 ns/op
BenchmarkCustomSha1_20000-2    	    1592	    749419 ns/op
BenchmarkBuiltinSha1_20000-2   	    1970	    589291 ns/op
PASS
ok  	github.com/danielost/sha-1	13.375s
```

The built-in implementation  is slightly faster as it uses additional optimizations.

**Benchmark results can be found in the `Actions` tab of the repository.**
