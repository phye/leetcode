package main

import "fmt"

type (
	Trie struct {
		Val    byte
		Child  *Trie
		Next   *Trie
		IsLeaf bool
	}
)

func Constructor() *Trie {
	trie := &Trie{' ', nil, nil, false}
	return trie
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	var p, n *Trie
	for c := this.Child; c != nil; p, c = c, c.Next {
		if c.Val == word[0] {
			n = c
			break
		}
	}
	if n == nil {
		n = &Trie{word[0], nil, nil, false}
		if p == nil {
			this.Child = n
		} else {
			p.Next = n
		}
	}
	n.Insert(word[1:])
	if len(word) == 1 {
		n.IsLeaf = true
	}
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		if this.Child == nil {
			return true
		} else {
			return this.IsLeaf
		}
	}
	var n *Trie
	for c := this.Child; c != nil; c = c.Next {
		if c.Val == word[0] {
			n = c
			break
		}
	}
	if n == nil {
		return false
	}
	return n.Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	var n *Trie
	for c := this.Child; c != nil; c = c.Next {
		if c.Val == prefix[0] {
			n = c
			break
		}
	}
	if n == nil {
		return false
	}
	return n.StartsWith(prefix[1:])
}

func main() {
	trie := Constructor()
	trie.Insert("to")
	trie.Insert("tea")
	trie.Insert("ted")
	trie.Insert("ten")
	trie.Insert("A")
	trie.Insert("i")
	trie.Insert("in")
	trie.Insert("inn")
	fmt.Printf("%v\n", trie)
	fmt.Printf("%v\n", trie.Search("ted"))
	fmt.Printf("%v\n", trie.Search("tex"))
}
