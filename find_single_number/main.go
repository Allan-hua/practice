package main

// Given an array of integers, arr,
// where all numbers occur twice except one number which occurs once, find the number.
// Your solution should ideally be O(n) time and use constant extra space.

/* bitmap */

func findSingle(nums []int) int {
	var bitmap uint64
	for _, n := range nums {
		b := uint64(n)
		bitmap = (bitmap | b) & ^(bitmap & b)
	}
	return int(bitmap)
}
