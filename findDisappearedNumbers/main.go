package main

import "fmt"

/*
Given an array of integers of size n,
where all elements are between 1 and n inclusive,
find all of the elements of [1, n] that do not appear in the array.
Some numbers may appear more than once.

Example:
Input: [4,5,2,6,8,2,1,5]
Output: [3,7]
class Solution(object):
  def findDisappearedNumbers(self, nums):
    # Fill this in.

nums = [4, 6, 2, 6, 7, 2, 1]
print(Solution().findDisappearedNumbers(nums))
# [3, 5]

For this problem, you can assume that you can mutate the input array.


*/
func main() {
	input := []int{4, 5, 2, 6, 8, 2, 1, 5}
	fmt.Println(findMissing1(input))
	fmt.Println(findMissing2(input))
}

func findMissing1(input []int) []int {
	var output []int
	cur := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		cur[input[i]-1] = input[i]
	}
	fmt.Println(cur)
	for i, n := range cur {
		if n == 0 {
			output = append(output, i+1)
		}
	}
	return output
}

func findMissing2(input []int) []int {
	var output []int
	for i := 0; i < len(input); i++ {

	}
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			output = append(output, i+1)
		}
	}
	return output
}
