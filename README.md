# Overview
A trie-based auto-completion system that allows efficient insertion, searching,
and prefix-based suggestions.

# Properties
- Use compressed trie by default
- Have support for standard trie
- Use utf-8 NFC to normalize utf-8 encoded string

# Example
```go
t := trie.New() // or trie.NewStandard for standard trie
t.Insert("car")
t.Insert("cart")
t.Search("car") // returns car, cart
```

# Implementation details
## Data structure
- compressed trie
- standard trie

## Algorithm
### Compressed trie
#### Insert
- insert text is empty -> return
- insert text is not empty
  - children is empty -> insert the word as a new child
  - children is not empty
    - loop through children, find the longest common prefix between a child & insert text
      - no found -> insert the word as a new children
      - some found, check the length
        - the prefix length == the word at node
          - the remaining of the inserting word == 0 -> set node's isWord = true
          - the remaining of the inserting word > 0 -> call word.insert with it
        - the prefix length < the word at node
          - the remaining of the inserting word == 0
            -> break the word at prefix, mark its isWord = true
            -> connect its children to the remaining part of it
            -> replace the node with the prefix node in its parent's children list
          - the remaining of the inserting word > 0
            -> break the word at prefix
            -> connect its children to the remaining part of it
            -> replace the node with the prefix node in its parent's children list
            -> insert the remaining of the inserting word as the new child

#### Search
- searchText is empty
  - node has no children
    - text so far is empty -> return
    - text so far is not empty -> collect
  - node has children
    -> collect text so far if node's isWord is true
    -> call child.collect, with text so far include text so far + child.value
- searchText is not empty
  - node has no children -> return
  - node has children
    - loop through children, find the longest common prefix
      - no found -> return
      - some found
        - both searchText & child value has remaining -> return
        - only searchText has remaining -> call child.search(text[prefixLength:])
        - only node value has remaining -> call child.search(text[prefixLength:])
        - no has remaining -> call child.search(text[prefixLength:])

### Standard trie
#### Insert
- insert text is empty -> return
- insert text is not empty
  -> break it into runes
  - node has no child
    -> insert runes[0] as a new child
    - runes[1:] is empty -> mark new child.isWord = true
    - runes[1:] is not empty -> call new child.insert(runes[1:])
  - node has child
    -> loop through the node children to see if any child has value as the first rune
     - no
       -> insert the rune as a new child
       - runes[1:] is empty -> mark new child.isWord = true
       - runes[1:] is not empty -> call new child.insert(runes[1:])
     - yes
       - runes[1:] is empty -> mark child.isWord = true
       - runes[1:] is not empty -> call new child.insert(runes[1:])

#### Search
- search text is empty
  - node has no children
    - text so far is empty -> return
    - text so far is not empty -> collect
  - node has children
    -> collect text so far if node's isWord is true
    -> call child.collect, with text so far include text so far + child.value
- search text is not empty
  - node has no children -> return
  - node has children
    -> find a child node with its value == runes[0]
       - no found, return
       - some found, call child.search(runes[1:], text so far + child.value)
