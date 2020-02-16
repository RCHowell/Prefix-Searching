package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleTrieInput = `to
top
cat
car`

var sampleTrie = &Node{
	Children: map[string]*Node{
		"t": &Node{
			Children: map[string]*Node{
				"o": &Node{
					IsWord: true,
					Children: map[string]*Node{
						"p": &Node{
							IsWord:   true,
							Children: map[string]*Node{},
						},
					},
				},
			},
		},
		"c": &Node{
			Children: map[string]*Node{
				"a": &Node{
					Children: map[string]*Node{
						"r": &Node{
							IsWord:   true,
							Children: map[string]*Node{},
						},
						"t": &Node{
							IsWord:   true,
							Children: map[string]*Node{},
						},
					},
				},
			},
		},
	},
}

func TestNewTrie(t *testing.T) {
	r := strings.NewReader(sampleTrieInput)
	constructedTrie := NewTrie(r)
	// b, _ := json.MarshalIndent(constructedTrie, "", "\t")
	// fmt.Println(string(b))
	// b, _ = json.MarshalIndent(sampleTrie, "", "\t")
	// fmt.Println(string(b))
	areEqual := reflect.DeepEqual(constructedTrie, sampleTrie)
	assert.Equal(t, areEqual, true)
}

func TestSearch(t *testing.T) {
	results := sampleTrie.Search("ca")
	assert.Equal(t, results, []string{"cat", "car"})
}
