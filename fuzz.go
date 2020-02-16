package main

import "math"

const assumedErrorRate = 0.2

// fuzzNode is useful for storing search metadata
type fuzzNode struct {
	n *Node
	l int
}

// fuzz is a map of character boundaries on a QWERTY keyboard
var fuzz = map[string][]string{
	"a": []string{"q", "w", "s", "z"},
	"b": []string{"v", "g", "h", "n"},
	"c": []string{"x", "d", "f", "v"},
	"d": []string{"s", "e", "r", "f", "c", "x"},
	"e": []string{"w", "s", "d", "r"},
	"f": []string{"e", "r", "t", "d", "g", "c", "v"},
	"g": []string{"t", "y", "h", "b", "v", "f"},
	"h": []string{"t", "y", "u", "j", "n", "b", "g"},
	"i": []string{"u", "j", "k", "o"},
	"j": []string{"u", "i", "k", "m", "n", "h"},
	"k": []string{"i", "o", "j", "l", "m"},
	"l": []string{"m", "p", "o"},
	"m": []string{"n", "j", "k"},
	"n": []string{"h", "j", "b", "m"},
	"o": []string{"i", "k", "l", "p"},
	"p": []string{"o", "l"},
	"q": []string{"a", "w"},
	"r": []string{"e", "d", "f", "t"},
	"s": []string{"a", "w", "e", "d", "x", "z"},
	"t": []string{"r", "f", "g", "h", "y"},
	"u": []string{"y", "h", "j", "k", "i"},
	"v": []string{"c", "f", "g", "b"},
	"w": []string{"q", "a", "s", "e"},
	"x": []string{"z", "s", "d", "c"},
	"y": []string{"t", "g", "h", "u"},
	"z": []string{"a", "s", "x"},
}

// FuzzSearch finds all words with the given prefix + searches down boundary paths
func (t *Node) FuzzSearch(prefix string) []string {

	results := make([]string, 0)
	stack := make([]*Node, 0)
	prefixLen := len(prefix)
	maxEditDistance := int(math.Round(float64(prefixLen) * assumedErrorRate))

	// Initialize fuzz stack
	fuzzStack := []*fuzzNode{
		&fuzzNode{
			n: t,
			l: 0,
		},
	}
	for len(fuzzStack) != 0 {
		n := len(fuzzStack) - 1
		node := fuzzStack[n]
		fuzzStack = fuzzStack[:n]

		// Each non-nil node at the end of the exploration is added to the stack
		if node.l == prefixLen {
			stack = append(stack, node.n)
			continue
		}
		// Let b(c) be the set of boundary characters for character c.
		// The search explores the characters {c} U b(c)
		ch := string(prefix[node.l])
		for _, c := range append(fuzz[ch], ch) {
			childNode := node.n.Children[c]
			childPrefix := node.n.acc + c
			if childNode != nil && editDistance(prefix, childPrefix) <= maxEditDistance {
				childNode.acc = childPrefix
				fuzzStack = append(fuzzStack, &fuzzNode{
					n: childNode,
					l: node.l + 1,
				})
			}
		}
	}

	for len(stack) != 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]
		if node.IsWord {
			results = append(results, node.acc)
		}
		for c, childNode := range node.Children {
			childNode.acc = node.acc + c
			stack = append(stack, childNode)
		}
	}

	return results
}

func editDistance(a, b string) int {
	var l, distance int
	la := len(a)
	lb := len(b)
	if la < lb {
		l = la
	} else {
		l = lb
	}
	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance
}
