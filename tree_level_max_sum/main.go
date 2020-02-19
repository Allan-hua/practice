package main

import "fmt"

// Node ...
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	n3 := Node{4, &Node{3, nil, nil}, &Node{2, nil, nil}}
	n2 := Node{5, &Node{4, nil, nil}, &Node{-1, nil, nil}}
	n1 := Node{1, &n2, &n3}
	fmt.Println(treeLevelMaxSum(n1))
}

func treeLevelMaxSum(root Node) (int, int) {
	var (
		level, max, count int
		heap              []*Node
	)
	heap = append(heap, root.Left, root.Right)
	count = len(heap)
	for len(heap) > 0 {
		var tmp []*Node
		var v, l int
		for count > 0 {
			if heap[0] != nil {
				tmp = append(tmp, heap[0])
			}
			heap = heap[1:]
			count--
		}
		l++
		for len(tmp) > 0 {
			if tmp[0].Left != nil {
				heap = append(heap, tmp[0].Left)
			}
			if tmp[0].Right != nil {
				heap = append(heap, tmp[0].Right)
			}
			v += tmp[0].Value
			tmp = tmp[1:]
			count++
		}
		if v > max {
			max = v
			level = l
		}
	}
	return level, max
}
