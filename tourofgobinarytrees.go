package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	fmt.Println("Tour of Go Solution - Equivalent Binary Trees\n")

	t1 := tree.New(5)
	t2 := tree.New(5)
	fmt.Println("Tree 1:", t1)
	fmt.Println("Tree 2:", t2)
	fmt.Println("Are they same? - ", Same(t1, t2))
	fmt.Println("------")

	t3 := tree.New(2)
	t4 := tree.New(3)
	fmt.Println("Tree 3:", t3)
	fmt.Println("Tree 4:", t4)
	fmt.Println("Are they same? - ", Same(t3, t4))
	fmt.Println("")
}

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	//close the channel to avoid panic
	close(ch)
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	} else {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	var br bool
	done := make(chan bool)
	c1 := make(chan int)
	c2 := make(chan int)
	go Walker(t1, c1)
	go Walker(t2, c2)
	go func() {
		for i := range c1 {
			if i == <-c2 {
				br = true
			} else {
				br = false
				break
			}
		}
		done <- true
	}()
	<-done
	return br
}
