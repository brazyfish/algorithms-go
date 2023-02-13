// huffmancoding_test.go
// description: Tests the compresssion, encoding and decoding algorithms of huffmancoding.go.
// author(s) [pedromsrocha](https://github.com/pedromsrocha)
// see huffmancoding.go

package huffman

import (
	"sort"
	"testing"
)

func SymbolCountOrd(message string) ByFreq {
	runeCount := make(map[rune]int)
	for _, s := range message {
		runeCount[s]++
	}
	listfreq := make(ByFreq, len(runeCount))
	i := 0
	for s, n := range runeCount {
		listfreq[i] = SymbolFreq{symbol: s, freq: n}
		i++
	}
	sort.Sort(listfreq)
	return listfreq
}

func TestHuffman(t *testing.T) {
	messages := []string{
		"hello world \U0001F600",
		"colorless green ideas sleep furiously",
		"the quick brown fox jumps over the lazy dog",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
		Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
		aliquip ex ea commodo consequat.`,
	}

	for _, message := range messages {
		t.Run("huffman: "+message, func(t *testing.T) {
			tree := BuildTree(SymbolCountOrd(message))
			dict := make(map[rune][]bool)
			BuildDict(&tree, nil, dict)
			for s, n := range dict {
				t.Logf("%q\t%v", s, n)
			}
			messageCoded := Encode(dict, message)
			messageDecoded := Decode(&tree, &tree, messageCoded, "")
			if messageDecoded != message {
				t.Errorf("got: %q\nbut expected: %q", messageDecoded, message)

			}
		})
	}
}
