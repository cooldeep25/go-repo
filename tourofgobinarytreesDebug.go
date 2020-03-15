package main

import (
	"fmt"
	"log"

	"golang.org/x/tour/tree"
)

func main() {

	log.Println("Tour of Go Solution")
	// Create a channel to synchronize goroutines
	done := make(chan bool)
	var si []int

	t := tree.New(1)
	log.Println("tree:", t)

	ch := make(chan int)
	t1 := tree.New(2)
	t2 := tree.New(1)

	fmt.Println("Tree 1:", t1)
	fmt.Println("Tree 2:", t2)

	func() {
		go Walker(t1, ch)
	}()

	fmt.Printf("Channel Print:\n")
	go func() {
		for i := range ch {
			// log.Println("Printing Value:", i)
			si = append(si, i)
		}
		log.Println("Channel Done")
		done <- true
	}()

	log.Println("Channels Complete:", <-done)
	log.Println("slice print:", si)

	// fmt.Printf("%v,", <-ch)
	// time.Sleep(5 * time.Second)
	// fmt.Println("\nAre they same? - ", Same(t1, t2))
	// time.Sleep(5 * time.Second)
}

func Walker(t *tree.Tree, ch chan int) {
	// Walk(t, ch)
	// time.Sleep(1 * time.Second)
	log.Println("Walkdebug")
	WalkDebug(t, ch)
	log.Println("Walk complete")
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

func WalkDebug(t *tree.Tree, ch chan int) {
	if t == nil {
		log.Println("empty channel")
		return
	} else if t.Left == nil {
		log.Println("Left Node Empty: Putting Current Node Value on Ch:", t.Value)
		ch <- t.Value
		if t.Right != nil {
			log.Println("Walking Right from node:", t.Value)
			Walk(t.Right, ch)
		}
		log.Println("Right Node also empty for node:", t.Value)
		log.Println("Returning walk from:", t.Value)
		log.Println("")
		return
	} else {
		log.Println("Walking Left from current node:", t.Value)
		Walk(t.Left, ch)
	}
	log.Println("Putting Value on Channel:", t.Value)
	ch <- t.Value
	if t.Right != nil {
		log.Println("Walking Right again from node: ", t.Value)
		Walk(t.Right, ch)
	}
	log.Println("End of Walk for node:", t.Value)
	log.Println("")
}

func Same(t1, t2 *tree.Tree) bool {
	var br bool
	log.Println("Bool Value:", br)
	done := make(chan bool)
	c1 := make(chan int)
	c2 := make(chan int)
	go Walker(t1, c1)
	go Walker(t2, c2)
	go func() {
		for i := range c1 {
			if i == <-c2 {
				// do nothing
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
