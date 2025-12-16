package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	// Expected
	exp := 4

	// Result
	res := count(b)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
