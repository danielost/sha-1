package sha1

import "math/rand"

// randSeq generates a random string of a specified length.
func randSeq(n int) string {
	// Define the character set to use for random string generation:
	var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&()*+,-./:;<=>?@[\\]^_`{|}~")
	b := make([]rune, n)

	// Generate random characters from the character set and append them to the buffer.
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
