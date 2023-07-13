package main

import "fmt"

type Node struct {
	value    int
	children []*Node
}

func BFS(root *Node, target int) bool {
	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.value == target {
			return true
		}

		for _, child := range node.children {
			queue = append(queue, child)
		}
	}

	return false
}

func main() {
	root := &Node{value: 1}
	node2 := &Node{value: 2}
	node3 := &Node{value: 3}
	node4 := &Node{value: 4}
	node5 := &Node{value: 5}

	root.children = []*Node{node2, node3}
	node2.children = []*Node{node4, node5}

	target := 5

	// BFS kullanarak hedef değeri arama
	found := BFS(root, target)

	if found {
		fmt.Printf("Hedef değer %d ağaçta bulundu.\n", target)
	} else {
		fmt.Printf("Hedef değer %d ağaçta bulunamadı.\n", target)
	}
}
