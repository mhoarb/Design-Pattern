package main

import "fmt"

var _ Component = &Folder{}

type Folder struct {
	component []Component
	name      string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("searching keyword %s in => %s\n ", keyword, f.name)
	for _, i := range f.component {
		i.search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.component = append(f.component, c)
}
