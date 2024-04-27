package main

import "fmt"

var _ Component = &file{}

type file struct {
	name string
}

func (f *file) Search(s string) {
	fmt.Printf("searching the %s keyword in %s\n", s, f.name)
}

func (f *file) getName() string {
	return f.name
}
