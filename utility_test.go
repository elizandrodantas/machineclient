package main

import "testing"

func TestUtility(t *testing.T) {
	random1 := randomByte(32)
	random2 := randomByte(16)

	if len(random1) != 32 {
		t.Errorf("Expected random byte to be of length %v, but got %v", 32, len(random1))
	}

	if len(random2) != 16 {
		t.Errorf("Expected random byte to be of length %v, but got %v", 16, len(random2))
	}
}
