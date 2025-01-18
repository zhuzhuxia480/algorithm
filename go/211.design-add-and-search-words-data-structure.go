/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */

// @lc code=start

package main
type WordDictionary struct {
    Word bool
	Next [26]*WordDictionary
}


func Constructor() WordDictionary {
    return WordDictionary{}
}


func (this *WordDictionary) AddWord(word string)  {
    curNode := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curNode.Next[index] == nil {
			newNode := &WordDictionary{}
			curNode.Next[index] = newNode
		}
		curNode = curNode.Next[index]
	}
	curNode.Word = true
}

func bfsNode(node *WordDictionary, word string) bool {
	q := []*WordDictionary{node}
	index := 0
	for len(q) > 0 && index < len(word) {
		qLen := len(q)
		flag := false
		for i := 0; i < qLen; i++ {
			curNode := q[i]
			nextNode := (*WordDictionary)(nil)
			if word[index] != '.' {
				ni := word[index] - 'a'
				nextNode = curNode.Next[ni]
				if nextNode != nil {
					if index == len(word)-1 && nextNode.Word{
						return true
					}
					q = append(q, nextNode)
					flag = true
				}
			} else {
				for j := 0; j < 26; j++ {
					nextNode = curNode.Next[j]
					if nextNode != nil {
						if index == len(word)-1 && nextNode.Word{
							return true
						}
						flag = true
						q = append(q, nextNode)
					}
				}
				
			}
			
		}
		if flag {
			index++
		}
		q = q[qLen:]
	}
	return false
}


func (this *WordDictionary) Search(word string) bool {
    return bfsNode(this, word)
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

