package trie

import (
	"slices"
	"testing"
)

func TestEmptyCompressedTrie(t *testing.T) {
	trie := newCompressed()
	if !slices.Equal(trie.collect(), []string{}) {
		t.Errorf("trie.collect() = %#v, want []string{}", trie.collect())
	}
}

func TestNotEmptyCompressedTrie(t *testing.T) {
	trie := newCompressed()
	trie.insert("car")
	trie.insert("cat")
	trie.insert("bar")
	if !slices.Equal(trie.collect(), []string{"car", "cat", "bar"}) {
		t.Errorf("trie.collect() = %#v, want []string{'car', 'cat', 'bar'}", trie.collect())
	}
}

func TestDuplicateCompressedTrieInsertion(t *testing.T) {
	trie := newCompressed()
	trie.insert("car")
	trie.insert("car")
	if !slices.Equal(trie.collect(), []string{"car"}) {
		t.Errorf("trie.collect() = %#v, want []string{'car'}", trie.collect())
	}
}

func TestEmbededCompressedTrieInsertion(t *testing.T) {
	trie := newCompressed()
	trie.insert("carpet")
	trie.insert("car")
	if !slices.Equal(trie.collect(), []string{"car", "carpet"}) {
		t.Errorf("trie.collect() = %#v, want []string{'car', 'carpet'}", trie.collect())
	}
}

func TestSeachNonExistValueCompressedTrie(t *testing.T) {
	trie := newCompressed()
	trie.insert("car")
	if !slices.Equal(trie.search("cat"), []string{}) {
		t.Errorf("trie.search('cat') = %#v, want []string{}", trie.search("cat"))
	}
}

func TestSeachExistValueCompressedTrie(t *testing.T) {
	trie := newCompressed()
	trie.insert("car")
	if !slices.Equal(trie.search("car"), []string{"car"}) {
		t.Errorf("trie.search('car') = %#v, want []string{'car'}", trie.search("car"))
	}
}

func TestPartialMatchCompressedTrie(t *testing.T) {
	trie := newCompressed()
	trie.insert("car")
	trie.insert("cat")
	if !slices.Equal(trie.search("ca"), []string{"car", "cat"}) {
		t.Errorf("trie.search('ca') = %#v, want []string{'car', 'cat'}", trie.search("ca"))
	}
}

func TestSearchEmptySearchQueryCompressedTrie(t *testing.T) {
	trie := newCompressed()
	trie.insert("car")
	if !slices.Equal(trie.search(""), []string{}) {
		t.Errorf("trie.search('') = %#v, want []string{}", trie.search(""))
	}
}

func TestSearchEmptyCompressedTrieAndEmptySearchQuery(t *testing.T) {
	trie := newCompressed()
	if !slices.Equal(trie.search(""), []string{}) {
		t.Errorf("trie.search('') = %#v, want []string{}", trie.search(""))
	}
}

func TestSearchEmptyCompressedTrie(t *testing.T) {
	trie := newCompressed()
	if !slices.Equal(trie.search("car"), []string{}) {
		t.Errorf("trie.search('car') = %#v, want []string{}", trie.search("car"))
	}
}
