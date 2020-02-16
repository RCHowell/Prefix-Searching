package main

import (
	"fmt"
	"log"
	"os"

	term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync() // cosmestic purpose
}

func main() {

	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	trie, ranks := setup()
	prefix := ""

	fmt.Println("> Prefix Searching")

keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyEnter:
				reset()
				prefix = ""
				fmt.Println(">")
			case term.KeyBackspace, term.KeyBackspace2, term.KeyDelete:
				reset()
				l := len(prefix)
				if l > 0 {
					prefix = prefix[:l-1]
				}
				searchAndPrint(trie, ranks, prefix)
			default:
				reset()
				prefix += string(ev.Ch)
				searchAndPrint(trie, ranks, prefix)
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}

func searchAndPrint(trie *Node, ranks map[string]int, prefix string) {
	if len(prefix) < 1 {
		return
	}
	fmt.Println(">", prefix)
	results := trie.FuzzSearch(prefix)
	filteredResults := CommonWords(ranks, results, 20)
	for _, w := range filteredResults {
		fmt.Println(w)
	}
}

// setup constructs the trie and rank map
func setup() (*Node, map[string]int) {
	wordsFile := "./internal/words.txt"
	r, err := os.Open(wordsFile)
	if err != nil {
		log.Fatal(err)
	}
	return NewTrie(r), GetWordRanks("./internal/ranks.json")
}
