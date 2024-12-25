/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
package main

type RandomizedSet struct {
	data []int
	index map[int]int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
		[]int{},
		map[int]int{},
	}
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.index[val]; ok {
		return false	
	}
	this.data = append(this.data, val)
	this.index[val] = len(this.data) - 1
	return true
}


func (this *RandomizedSet) Remove(val int) bool {
    id, ok := this.index[val]
	if !ok {
		return false
	}

	last := len(this.data) - 1

	this.data[id] = this.data[last]
	this.index[this.data[last]] = id
	this.data = this.data[:last]
	delete(this.index, val)
	return true
}


func (this *RandomizedSet) GetRandom() int {
    return this.data[rand.Intn(len(this.data))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

