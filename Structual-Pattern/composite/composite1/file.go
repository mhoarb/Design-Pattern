package main

import "fmt"

var _ Component = &File{}

type File struct {
	name string
}

func (f *File) search(s string) {
	fmt.Printf("search the keyword %s in =>%s\n", s, f.name)
}

func (f *File) getName() string {
	return f.name
}
