package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ip := generator("1592551013", []string{})
	fmt.Println(strings.Join(ip, "."))
}

func generator(s string, parts []string) (res []string) {
	if len(parts) == 4 {
		if s == "" {
			return parts
		}
		return []string{}
	}
	str := strings.Split(s, "")
	// too long
	if len(str) > 3*(4-len(parts)) {
		return []string{}
	}
	// too short
	if len(str) < 4-len(parts) {
		return []string{}
	}
	if str[0] == "0" {
		return generator(strings.Join(str[1:], ""), append(parts, str[0]))
	}

	l := 3
	if len(str) < l {
		l = len(str)
	}

	for i := 0; i < l; i++ {
		tmp := strings.Join(str[:i], "")
		s = strings.Join(str[i:], "")
		if n, _ := strconv.Atoi(tmp); n < 256 {
			parts = append(parts, tmp)
			fmt.Println(len(strings.Split(s, "")), len(parts), s, parts, tmp, res)
			res = append(res, generator(s, parts)...)
		}
	}

	return res
}
