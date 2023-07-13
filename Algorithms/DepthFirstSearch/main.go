package main

import "fmt"

type Node struct {
	value    int
	visited  bool
	children []*Node
}

func DFS(node *Node, target int) bool {
	if node.value == target {
		return true
	}

	node.visited = true

	for _, child := range node.children {
		if !child.visited {
			found := DFS(child, target)
			if found {
				return true
			}
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

	// DFS kullanarak hedef değeri arama
	found := DFS(root, target)

	if found {
		fmt.Printf("Hedef değer %d ağaçta bulundu.\n", target)
	} else {
		fmt.Printf("Hedef değer %d ağaçta bulunamadı.\n", target)
	}
}
