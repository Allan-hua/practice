package main

import "fmt"

// Starting at index 0, for an element n at index i,
// you are allowed to jump at most n indexes ahead. Given a list of numbers,
// find the minimum number of jumps to reach the end of the list.

type arr []int

func jumpToEnd(nums []int) {
	var a arr
	a = nums
	fmt.Println(a.jump(0, 0))
}

func (a arr) jump(index, count int) int {
	count++
	var max int
	if index+a[index] > len(a)-1 {
		return count
	}
	for i := index; i < index+a[index]; i++ {
		if a[i]+i-index > max {
			max = a[i] + i - index
		}
	}
	return a.jump(index+max, count)
}
