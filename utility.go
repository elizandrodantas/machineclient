package main

import "crypto/rand"

// randomByte generates a slice of random bytes with the specified length.
func randomByte(len int) []byte {
	buf := make([]byte, len)
	rand.Read(buf)
	return buf
}
