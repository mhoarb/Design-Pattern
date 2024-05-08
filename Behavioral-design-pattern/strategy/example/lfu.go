package main

import "fmt"

type Lfu struct {
}

func (l *Lfu) evict(c *Cash) {
	fmt.Println("evicting by lfu strtegy")
}
