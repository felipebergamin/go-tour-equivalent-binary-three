package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go Walk(t1, chan1)
	go Walk(t2, chan2)

	// three.New(k) always returns a three with 10 elements: k, 2k, 3k, ..., 10k
	for i := 0; i < 10; i += 1 {
		node1, node2 := <-chan1, <-chan2
		fmt.Println(node1, node2)
		if node1 != node2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Is equivalent? ", Same(tree.New(4), tree.New(4)))
	//c := make(chan int)
	//go Walk(tree.New(5), c)
	//
	//for node1 := range c {
	//	fmt.Println(node1)
	//}
}
