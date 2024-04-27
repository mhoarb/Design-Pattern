package main

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		component: nil,
		name:      "Folder1",
	}
	folder1.Add(file1)

	folder2 := &Folder{
		component: nil,
		name:      "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)

	folder2.search("rose")

}
