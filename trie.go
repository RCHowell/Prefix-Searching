package main

import (
	"bufio"
	"io"
)

// Node is the core type for a trie
type Node struct {
	Children map[string]*Node
	IsWord   bool
	acc      string
}

// NewTrie constructs a prefix tree from a list of words
func NewTrie(r io.Reader) *Node {

	root := &Node{
		Children: make(map[string]*Node),
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		word := s.Text()
		curr := root
		for _, c := range word {
			n := curr.Children[string(c)]
			if n == nil {
				n = &Node{
					Children: make(map[string]*Node),
				}
				curr.Children[string(c)] = n
			}
			curr = n
		}
		curr.IsWord = true
	}

	return root
}

// Search finds all words in the prefix's sub-trie (limiting to <= n results)
func (t *Node) Search(prefix string) []string {

	results := make([]string, 0)
	curr := t

	// Traverse trie until the end of prefix
	// Terminate early if no prefix match in the trie
	for _, c := range prefix {
		node := curr.Children[string(c)]
		if node == nil {
			return []string{"No Results"}
		}
		curr = node
	}

	// Search for words in the sub-trie
	curr.acc = prefix
	stack := []*Node{curr}
	for len(stack) != 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]
		if node.IsWord {
			results = append(results, node.acc)
		}
		for k, v := range node.Children {
			v.acc = node.acc + k
			stack = append(stack, v)
		}
	}

	return results
}
