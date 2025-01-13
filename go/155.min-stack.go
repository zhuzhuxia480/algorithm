/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
package main
type MinStack struct {
    data [][]int
}


func Constructor() MinStack {
    return MinStack{
		data: make([][]int, 0),
	}
}


func (this *MinStack) Push(val int)  {
	currentMin := 0
    if len(this.data) == 0 {
		currentMin = val
	} else {
		currentMin = min(val, this.GetMin())
	}
	this.data = append(this.data, []int{val, currentMin})
}


func (this *MinStack) Pop()  {
    this.data = this.data[:len(this.data)-1]
}


func (this *MinStack) Top() int {
    return this.data[len(this.data)-1][0]
}


func (this *MinStack) GetMin() int {
    return this.data[len(this.data)-1][1]
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

