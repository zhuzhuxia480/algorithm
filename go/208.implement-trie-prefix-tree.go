/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
package main

type Node struct {
	Val byte
	Word bool
	Next [26]*Node
}

type Trie struct {
    HeadNode Node
}


func Constructor() Trie {
    return Trie{
		HeadNode: Node{
			Val: '-',
			Word: false,
		},
	}
}


func (this *Trie) Insert(word string)  {
    curNode := &this.HeadNode
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curNode.Next[index] == nil {
			nexNode := &Node{
				Val: word[i],
				Word: i == len(word)-1,
			}
			curNode.Next[index] = nexNode
		} else {
			if i == len(word) - 1 {
				curNode.Next[index].Word = true
			}
		}
		curNode = curNode.Next[index]
	}
}


func (this *Trie) Search(word string) bool {
    curNode := &this.HeadNode
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curNode.Next[index] == nil {
			return false
		}
		if i == len(word)-1 && curNode.Next[index].Word {
			return true
		}
		curNode = curNode.Next[index]
	}
	return false
}


func (this *Trie) StartsWith(prefix string) bool {
    curNode := &this.HeadNode
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if curNode.Next[index] == nil {
			return false
		}
		if i == len(prefix)-1 {
			return true
		}
		curNode = curNode.Next[index]
	}
	return false
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

