package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuzzSearch(t *testing.T) {
	expected := []string{"car", "cat"}
	// "caf" could be a typo for "car" or "cat"
	assert.Equal(t, expected, sampleTrie.FuzzSearch("caf"))
}
