package main

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

func main() {
	t := tree.New(1)
	// for a binary tree -
	// Left Node Value  < Node Value < Right Node Value
	ch := make(chan int)
	// Walk the Tree and Store the values on Channel
	go Walk(t, ch)
	// Print the Channel Values stored in above goroutine
	go func() {
		fmt.Printf("Channel Print:")
		for i := range ch {
			fmt.Printf("%v,", i)
		}
	}()
	//Check whether the Two Trees stores the same values.
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println("\nAre they same? - ", Same(t1, t2))
	time.Sleep(1 * time.Second)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// check if there is No Node on the Left side then return same node
	// If there is Left Node present return the Left Node to Traverse Further
	if t == nil {
		return
	} else if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	} else {
		// Further Traverse the Binary Tree to Get Left Node
		// This is recursive function call
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var br bool
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	go func() {
		// Access the Channels Elements and Compare them
		for i := range c1 {
			if i == <-c2 {
				br = true
			} else {
				br = false
			}
		}
	}()
	time.Sleep(100 * time.Millisecond)
	return br
}
