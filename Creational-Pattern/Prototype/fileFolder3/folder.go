package main

import "fmt"

var _ INode = &Folder{}

type Folder struct {
	children []INode
	name     string
}

func (f Folder) print(indaentation string) {
	fmt.Println(indaentation + f.name)
	for _, i := range f.children {
		i.print(" ")
	}
}

func (f Folder) clone() INode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []INode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
