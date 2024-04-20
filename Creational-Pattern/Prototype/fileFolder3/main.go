package main

import "fmt"

func main() {
	file1 := File{name: "file1"}
	file2 := File{name: "file2"}
	file3 := File{name: "file3"}

	folder1 := &Folder{
		children: []INode{file1},
		name:     "Folder1",
	}
	folder2 := &Folder{
		children: []INode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("printing folder 1")
	folder1.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("1printing for foldr 2")

	cloneFolder.print("  ")
}
