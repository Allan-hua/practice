package main

import (
	"fmt"
	"strings"
)

/*
Given a list of words in a string, reverse the words in-place (ie don’t create a new string and reverse the words). Since python strings are not mutable, you can assume the input will be a mutable sequence (like list).
Here’s an example and some starting code:

def reverse_words(words):
  # Fill this in.
s = list("can you read this")
reverse_words(s)
print(''.join(s))
# this read you can
*/
func main() {
	str := strings.Split("can you read this", "")
	fmt.Println(reverseWords(str))
}

func reverseWords(str []string) string {
	for i := 0; i < len(str); i++ {
		if str[i] == " " {
			for j := i + 1; j < len(str); j++ {
				if j == len(str)-1 {
					shift(str, i)
					str[0] = " "
					for i != len(str)-1 {
						shift(str, len(str)-1)
						i++
					}
					break
				}
				if str[j] == " " {
					for i != j {
						shift(str, j)
						i++
					}
				}
			}
		}
	}
	res := strings.Join(str, "")
	return res
}

func shift(str []string, k int) []string {
	tmp := str[k]
	for k > 0 {
		str[k] = str[k-1]
		k--
	}
	str[0] = tmp
	return str
}
