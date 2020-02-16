package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWordRanks(t *testing.T) {
	ranks := GetWordRanks("./internal/ranks.json")
	assert.Equal(t, ranks["dog"], 1047)
	assert.Equal(t, ranks["turtle"], 9688)
	assert.Equal(t, ranks["australopithecus"], 149259)
}

func TestCommonWords(t *testing.T) {
	ranks := GetWordRanks("./internal/ranks.json")
	words := []string{
		"green",
		"the",
		"turtle",
	}
	expectedWordsInOrder := []string{"the", "green", "turtle"}
	expectedTwoWords := []string{"the", "green"}
	assert.Equal(t, CommonWords(ranks, words, 3), expectedWordsInOrder)
	assert.Equal(t, CommonWords(ranks, words, 2), expectedTwoWords)
}
