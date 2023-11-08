// Package sha1 implements the SHA-1 hash algorithm.
//
// The implementation is based on the pseudo code from the Wikipedia article:
// https://en.wikipedia.org/wiki/SHA-1
package sha1

import (
	"encoding/binary"
	"math/big"
)

// Sum calculates the SHA-1 hash of the given message.
func Sum(message []byte) []byte {
	// Initialize variables:
	var (
		h0 uint32 = 0x67452301
		h1 uint32 = 0xEFCDAB89
		h2 uint32 = 0x98BADCFE
		h3 uint32 = 0x10325476
		h4 uint32 = 0xC3D2E1F0
	)

	// Pad the message to be a multiple of 512 bits:
	addPadding(&message)

	// Process the message in successive 512-bit chunks:
	chunks := breakMessageIntoChunks(&message, 512)

	for _, chunk := range chunks {
		// Break each chunk into 32-bit words:
		wordsByte := breakMessageIntoChunks(&chunk, 32)
		// Allocate an array to hold 80 uint32 values:
		wordsUint := make([]uint32, 80)

		// Convert each byte to an uint32:
		for i, b := range wordsByte {
			wordsUint[i] = binary.BigEndian.Uint32(b)
		}

		// Message schedule: extend the sixteen 32-bit words into eighty 32-bit words:
		for i := 16; i <= 79; i++ {
			wordsUint[i] = leftRotate(wordsUint[i-3]^wordsUint[i-8]^wordsUint[i-14]^wordsUint[i-16], 1)
		}

		// Initialize hash value for this chunk:
		a := h0
		b := h1
		c := h2
		d := h3
		e := h4

		// Main Loop:
		var f, k uint32
		for i := 0; i <= 79; i++ {
			if i <= 19 {
				f = (b & c) | ((^b) & d)
				k = 0x5A827999
			} else if i <= 39 {
				f = b ^ c ^ d
				k = 0x6ED9EBA1
			} else if i <= 59 {
				f = (b & c) | (b & d) | (c & d)
				k = 0x8F1BBCDC
			} else {
				f = b ^ c ^ d
				k = 0xCA62C1D6
			}

			// Calculate the cureent intermediate value:
			temp := leftRotate(a, 5) + f + e + k + wordsUint[i]
			e = d
			d = c
			c = leftRotate(b, 30)
			b = a
			a = temp
		}

		// Add this chunk's hash to result so far:
		h0 = h0 + a
		h1 = h1 + b
		h2 = h2 + c
		h3 = h3 + d
		h4 = h4 + e
	}

	// Produce the final hash value (big-endian) as a 160-bit number:
	hh0 := big.NewInt(int64(h0))
	hh1 := big.NewInt(int64(h1))
	hh2 := big.NewInt(int64(h2))
	hh3 := big.NewInt(int64(h3))
	hh4 := big.NewInt(int64(h4))

	hh0 = hh0.Lsh(hh0, 128)
	hh1 = hh1.Lsh(hh1, 96)
	hh2 = hh2.Lsh(hh2, 64)
	hh3 = hh3.Lsh(hh3, 32)

	hhs := []*big.Int{hh0, hh1, hh2, hh3, hh4}
	hh := big.NewInt(0)
	for _, currHh := range hhs {
		hh = hh.Or(hh, currHh)
	}
	bytes := hh.Bytes()
	for len(bytes) < 20 {
		bytes = append([]byte{0x00}, bytes...)
	}
	return bytes
}

// addPadding pads the message to ensure its length is a multiple of 512 bits.
func addPadding(message *[]byte) {
	messageLength := uint64(len(*message) * 8)
	*message = append(*message, 0x80)

	// Pad the message with 0x00 bytes until the length is a multiple of 512 bits.
	// Leaves space for the message length.
	for len(*message)%64 != 56 {
		*message = append(*message, 0x00)
	}

	// Append the message length:
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, messageLength)
	*message = append(*message, b...)
}

// breakMessageIntoChunks breaks the message into N-bit chunks.
func breakMessageIntoChunks(message *[]byte, chunkSize int) [][]byte {
	// Calculate the number of chunks to split the message into:
	chunkSize = chunkSize / 8
	numberOfChunks := len(*message) / chunkSize

	// Create an empty slice to store the chunks.
	chunks := make([][]byte, numberOfChunks)

	// Split the message into chunks.
	for i := 0; i < numberOfChunks; i++ {
		chunks[i] = (*message)[i*chunkSize : i*chunkSize+chunkSize]
	}
	return chunks
}

// leftRotate rotates the 32-bit unsigned integer u by n bits to the left.
func leftRotate(u uint32, n int) uint32 {
	return (u << n) | (u >> (32 - n))
}
