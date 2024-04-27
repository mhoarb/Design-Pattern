package main

import "fmt"

var _ INode = &file{}

type file struct {
	name string
}

func (f file) print(identitation string) {
	fmt.Println(identitation + f.name)
}

func (f file) clone() INode {
	return &file{name: f.name + "_clone"}
}
