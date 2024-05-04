package main

import (
	"fmt"
)

type Tv struct {
	name     string
	isRunnig bool
}

func (t *Tv) on() {
	if t.isRunnig == true {
		fmt.Printf("this tv name : %s is already ON\n", t.name)
	}
	fmt.Printf("turning on %s tv\n", t.name)

	t.isRunnig = true
	fmt.Printf("%s is ON\n", t.name)
}

func (t *Tv) off() {
	if t.isRunnig == false {
		fmt.Printf("this tv %s is already off\n", t.name)
	}
	fmt.Printf("turning tv %s off\n", t.name)

	t.isRunnig = false
	fmt.Println("turning off this tv")
}
