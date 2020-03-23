package main

import "fmt"

/*
Hi, here's your problem today. This problem was recently asked by Twitter:
Given 3 sorted lists, find the intersection of those 3 lists.
Here's an example and some starter code.
def intersection(list1, list2, list3):
  # Fill this in.
print(intersection([1, 2, 3, 4], [2, 4, 6, 8], [3, 4, 5]))
# [4]
*/

func main() {
	list1 := []int{1, 2, 3, 4}
	list2 := []int{2, 4, 6, 8}
	list3 := []int{3, 4, 5}
	fmt.Println(intersection(list1, list2, list3))
}

func intersection(list1, list2, list3 []int) (res []int) {
	in := make([]int, len(list1)+len(list2)+len(list3))
	for _, num := range list1 {
		in[num]++
	}
	for _, num := range list2 {
		in[num]++
	}
	for _, num := range list3 {
		in[num]++
	}
	for i, n := range in {
		if n == 3 {
			res = append(res, i)
		}
	}
	return res
}
