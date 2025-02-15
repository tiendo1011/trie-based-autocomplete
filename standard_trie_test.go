package trie

import (
	"slices"
	"testing"
)

func TestEmptyTrie(t *testing.T) {
	trie := newStandard()
	if !slices.Equal(trie.collect(), []string{}) {
		t.Errorf("trie.collect() = %#v, want []string{}", trie.collect())
	}
}

func TestNotEmptyTrie(t *testing.T) {
	trie := newStandard()
	trie.insert("car")
	trie.insert("cat")
	trie.insert("bar")
	if !slices.Equal(trie.collect(), []string{"car", "cat", "bar"}) {
		t.Errorf("trie.collect() = %#v, want []string{'car', 'cat', 'bar'}", trie.collect())
	}
}

func TestDuplicateInsertion(t *testing.T) {
	trie := newStandard()
	trie.insert("car")
	trie.insert("car")
	if !slices.Equal(trie.collect(), []string{"car"}) {
		t.Errorf("trie.collect() = %#v, want []string{'car'}", trie.collect())
	}
}

func TestEmbededInsertion(t *testing.T) {
	trie := newStandard()
	trie.insert("car")
	trie.insert("carpet")
	if !slices.Equal(trie.collect(), []string{"car", "carpet"}) {
		t.Errorf("trie.collect() = %#v, want []string{'car', 'carpet'}", trie.collect())
	}
}

func TestSeachNonExistValue(t *testing.T) {
	trie := newStandard()
	trie.insert("car")
	if !slices.Equal(trie.search("cat"), []string{}) {
		t.Errorf("trie.search('cat') = %#v, want []string{}", trie.search("cat"))
	}
}

func TestSeachExistValue(t *testing.T) {
	trie := newStandard()
	trie.insert("car")
	if !slices.Equal(trie.search("car"), []string{"car"}) {
		t.Errorf("trie.search('car') = %#v, want []string{'car'}", trie.search("car"))
	}
}

func TestPartialMatch(t *testing.T) {
	trie := newStandard()
	trie.insert("car")
	trie.insert("cat")
	if !slices.Equal(trie.search("ca"), []string{"car", "cat"}) {
		t.Errorf("trie.search('ca') = %#v, want []string{'car', 'cat'}", trie.search("ca"))
	}
}

func TestSearchEmptySearchQuery(t *testing.T) {
	trie := newStandard()
	trie.insert("car")
	if !slices.Equal(trie.search(""), []string{}) {
		t.Errorf("trie.search('') = %#v, want []string{}", trie.search(""))
	}
}

func TestSearchEmptyTrieAndEmptySearchQuery(t *testing.T) {
	trie := newStandard()
	if !slices.Equal(trie.search(""), []string{}) {
		t.Errorf("trie.search('') = %#v, want []string{}", trie.search(""))
	}
}

func TestSearchEmptyTrie(t *testing.T) {
	trie := newStandard()
	if !slices.Equal(trie.search("car"), []string{}) {
		t.Errorf("trie.search('car') = %#v, want []string{}", trie.search("car"))
	}
}
