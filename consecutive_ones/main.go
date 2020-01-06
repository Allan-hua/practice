package main

import "fmt"

/*
Return the longest run of 1s for a given integer n's binary representation.

Example:
Input: 242
Output: 4
242 in binary is 0b11110010, so the longest run of 1 is 4.

def longest_run(n):
  # Fill this in.

print longest_run(242)
# 4
*/

func main() {
	fmt.Println(logestRun(242))
}

func logestRun(num int) int {
	b := uint64(num)
	var max, tmp int
	for i := 0; i < 64; i++ {
		if b&(1<<uint(i)) != 0 {
			tmp++
		} else if tmp > max {
			max = tmp
		} else {
			tmp = 0
		}
	}
	return max
}
