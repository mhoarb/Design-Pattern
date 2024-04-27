package main

import "fmt"

var _ Component = &Folder{}

type Folder struct {
	component []Component
	name      string
}

func (f *Folder) Search(s string) {
	fmt.Printf("searching the keyword : %s in this =>%s\n", s, f.name)
	for _, i := range f.component {
		i.Search(s)
	}

}

func (f *Folder) Add(component Component) {
	f.component = append(f.component, component)
}
