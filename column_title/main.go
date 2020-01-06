package main

import (
	"fmt"
	"strings"
)

/*
MS Excel column titles have the following pattern: A, B, C, ..., Z, AA, AB, ..., AZ, BA, BB, ..., ZZ, AAA, AAB, ... etc. In other words, column 1 is named "A", column 2 is "B", column 26 is "Z", column 27 is "AA" and so forth. Given a positive integer, find its corresponding column name.
Examples:
Input: 26
Output: Z

Input: 51
Output: AY

Input: 52
Output: AZ

Here is a starting point:

class Solution:
    def convertToTitle(self, n):
        from itertools import product, islice
        def gen():
            r = 1
            while True:
                for i in product('ABCDEFGHIJKLMNOPQRSTUVWXYZ', repeat=r):
                    yield ''.join(i)
                r+=1
        return list(islice(gen(),n))[-1]

input1 = 1
input2 = 456976
input3 = 28
print(Solution().convertToTitle(input1))
# A
print(Solution().convertToTitle(input2))
# YYYZ
print(Solution().convertToTitle(input3))
# AB
*/

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(456976))
	fmt.Println(convertToTitle(28))
}

func convertToTitle(input int) string {
	a := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	var tmp, title []string
	for {
		if input%26 > 0 {
			tmp = append(tmp, a[input%26-1])
			input -= input % 26
		}
		if input/26 < 26 && input > 0 {
			tmp = append(tmp, a[input/26-1])
			input = input / 26
		} else {
			tmp = 
		}
		if input == 0 {
			break
		}
	}
	if len(tmp) == 1 {
		return strings.Join(tmp, "")
	}
	for i := len(tmp); i > 0; i-- {
		title = append(title, tmp[i-1])
	}
	return strings.Join(title, "")
}
