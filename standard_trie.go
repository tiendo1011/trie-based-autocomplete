package trie

import (
	"golang.org/x/text/unicode/norm"
)

type node struct {
	value    rune
	children map[rune]*node
	isWord   bool
}

type trie struct {
	root *node
}

func newStandard() *trie {
	return &trie{root: &node{children: make(map[rune]*node)}}
}

func (t *trie) insert(text string) {
	if text == "" {
		return
	}

	normalizedText := norm.NFC.String(text)
	runes := []rune(normalizedText)
	t.root.insert(runes)
}

func (n *node) insert(runes []rune) {
	first := runes[0]
	matching_child, ok := n.children[first]
	if !ok {
		matching_child = &node{value: first, children: make(map[rune]*node)}
		n.children[first] = matching_child
	}
	if len(runes[1:]) == 0 {
		matching_child.isWord = true
	} else {
		matching_child.insert(runes[1:])
	}
}

func (t *trie) search(text string) []string {
	if text == "" {
		return []string{}
	}

	normalizedText := norm.NFC.String(text)
	runes := []rune(normalizedText)

	result := []string{}
	t.root.search("", runes, &result)

	return result
}

func (n *node) search(text string, runes []rune, result *[]string) {
	if len(runes) == 0 {
		n.collect(text, result)
		return
	}

	if len(n.children) == 0 {
		return
	}

	if child, ok := n.children[runes[0]]; ok {
		child.search(text+string(child.value), runes[1:], result)
	}
}

func (t *trie) collect() []string {
	result := []string{}
	t.root.collect("", &result)

	return result
}

func (n *node) collect(text string, result *[]string) {
	if len(n.children) == 0 {
		if text != "" {
			*result = append(*result, text)
		}

		return
	}

	for _, child := range n.children {
		if n.isWord {
			*result = append(*result, text)
		}
		child.collect(text+string(child.value), result)
	}
}
