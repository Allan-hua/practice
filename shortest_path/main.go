package main

import (
	"fmt"
	"strings"
)

/*
Input: '/Users/Joma/Documents/../Desktop/./../'
Output: '/Users/Joma/'
def shortest_path(file_path):
  # Fill this in.

print shortest_path('/Users/Joma/Documents/../Desktop/./../')
# /Users/Joma/
*/

func main() {
	fmt.Println(shortestPath("/Users/Joma/Documents/../Desktop/./../"))
}

func shortestPath(path string) string {
	strs := strings.Split(path, "/")
	var nPath []string
	for _, s := range strs {
		if s == "." {
			continue
		}
		if s == ".." {
			nPath = nPath[:len(nPath)-1]
			continue
		}
		nPath = append(nPath, s)
	}
	fmt.Println(strs)
	return strings.Join(nPath, "/")
}
