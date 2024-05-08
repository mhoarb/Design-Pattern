package main

import "fmt"

type Fifo struct {
}

func (f *Fifo) evict(c *Cash) {
	fmt.Println("evicting by fifo strategy")
}
