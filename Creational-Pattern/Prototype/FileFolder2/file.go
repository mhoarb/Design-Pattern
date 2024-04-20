package main

import "fmt"

var _ INode = &file{}

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() INode {
	return &file{name: f.name + "_clone"}
}
