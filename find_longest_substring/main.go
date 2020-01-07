package main

import (
	"fmt"
	"strings"
)

func main() {
	findLongestString("pwwkew")

}

func findLongestString(input string) {
	var max, start int
	var str []string
	var res string
	in := strings.Split(input, "")
	tmp := make([]int, 26)
	for i := 0; i < len(in); i++ {
		fmt.Print(in[i])
		str = append(str, in[i])
		for n := 0; n < tmp[toNum(in[i])]-start+1; n++ {
			tmp[toNum(str[n])] = 0
			str = str[1:]
			start++
		}
		tmp[toNum(in[i])] = i
		if max < i-start+1 {
			max = i - start + 1
			res = strings.Join(str, "")
		}
	}
	fmt.Println(max, res)
}

func toNum(s string) int {
	r := rune(([]rune)(s)[0])
	return int(r - 'a')
}