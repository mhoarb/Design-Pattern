package main

import "fmt"

type Lru struct {
}

func (l *Lru) evict(c *Cash) {
	fmt.Println("Evicting by fifo strategy")
}
