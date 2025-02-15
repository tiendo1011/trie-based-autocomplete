package trie

import (
	"golang.org/x/text/unicode/norm"
)

type compressedNode struct {
	value    string
	isWord   bool
	children []*compressedNode
}

type compressedTrie struct {
	root *compressedNode
}

func newCompressed() *compressedTrie {
	return &compressedTrie{root: &compressedNode{}}
}

func (t *compressedTrie) insert(text string) {
	t.root.insert(norm.NFC.String(text))
}

func (n *compressedNode) insert(text string) {
	if text == "" {
		return
	}

	if len(n.children) == 0 {
		newNode := &compressedNode{value: text, isWord: true}
		n.children = append(n.children, newNode)
		return
	}

	for i, child := range n.children {
		if prefixLength := child.commonPrefixLength(text); prefixLength != 0 {
			if prefixLength == len(child.value) {
				if len(text[prefixLength:]) == 0 {
					child.isWord = true
				} else {
					child.insert(text[prefixLength:])
				}
			} else if prefixLength < len(child.value) {
				prefixNode := &compressedNode{value: child.value[:prefixLength]}
				n.children[i] = prefixNode

				remaining := &compressedNode{value: child.value[prefixLength:]}
				remaining.children = child.children

				prefixNode.children = append(prefixNode.children, remaining)

				if len(text[prefixLength:]) == 0 {
					prefixNode.isWord = true
				} else {
					insertNode := &compressedNode{value: text[prefixLength:], isWord: true}
					prefixNode.children = append(prefixNode.children, insertNode)
				}
			} else {
				panic("len(prefix) > len(child.value)")
			}
			return
		}
	}

	// no common prefix found
	newNode := &compressedNode{value: text, isWord: true}
	n.children = append(n.children, newNode)
}

func (n *compressedNode) commonPrefixLength(text string) int {
	for i := range n.value {
		if i >= len(text) || n.value[i] != text[i] {
			return i
		}
	}

	return len(n.value)
}

func (t *compressedTrie) search(text string) []string {
	if text == "" {
		return []string{}
	}

	var result []string
	t.root.search("", norm.NFC.String(text), &result)

	return result
}

func (n *compressedNode) search(textSoFar string, searchText string, result *[]string) {
	if searchText == "" {
		n.collect(textSoFar, result)
		return
	}

	for _, child := range n.children {
		if prefixLength := child.commonPrefixLength(searchText); prefixLength != 0 {
			if len(child.value[prefixLength:]) > 0 && len(searchText[prefixLength:]) > 0 {
				continue
			}

			child.search(textSoFar+child.value, searchText[prefixLength:], result)
		}
	}
}

func (t *compressedTrie) collect() []string {
	result := []string{}
	t.root.collect("", &result)

	return result
}

func (n *compressedNode) collect(text string, result *[]string) {
	if len(n.children) == 0 {
		if text != "" {
			*result = append(*result, text)
		}

		return
	}

	if n.isWord {
		*result = append(*result, text)
	}

	for _, child := range n.children {
		child.collect(text+child.value, result)
	}
}
