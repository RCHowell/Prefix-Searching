package main

import (
	"container/heap"
	"encoding/json"
	"io/ioutil"
)

type word struct {
	text string
	rank int
}

type wordHeap []*word

func (h wordHeap) Len() int           { return len(h) }
func (h wordHeap) Less(i, j int) bool { return h[i].rank < h[j].rank }
func (h wordHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *wordHeap) Push(x interface{}) {
	*h = append(*h, x.(*word))
}

func (h *wordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// GetWordRanks reads the internal frequency map into memory
func GetWordRanks(filePath string) map[string]int {
	ranks := make(map[string]int)
	bytes, _ := ioutil.ReadFile(filePath)
	_ = json.Unmarshal(bytes, &ranks)
	return ranks
}

// CommonWords takes a list of words and returns the n most common
// If len(words) < n, then it will return the len(words) most common words
func CommonWords(ranks map[string]int, words []string, n int) []string {

	h := &wordHeap{}
	heap.Init(h)

	for _, w := range words {
		heap.Push(h, &word{
			text: w,
			rank: ranks[w],
		})
	}

	heapLen := h.Len()
	if heapLen < n {
		n = heapLen
	}

	results := make([]string, n)
	for i := 0; i < n; i++ {
		results[i] = (*h)[i].text
	}

	return results
}
