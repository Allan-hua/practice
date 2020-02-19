package main

import "fmt"

func main() {
	nums := []int{-5, -3, -1, 0, 1, 4, 5}
	fmt.Println(squareNumbers(nums))
}

func squareNumbers(nums []int) (sorted []int) {
	tmp := make([]int, len(nums))
	for len(nums) > 0 {
		n := 0
		if nums[0] < 0 {
			n = -nums[0]
		} else {
			n = nums[0]
		}
		tmp[n]++
		nums = nums[1:]
	}
	for i, v := range tmp {
		for j := 0; j < v; j++ {
			sorted = append(sorted, i*i)
		}
		tmp = tmp[1:]
	}
	return sorted
}
