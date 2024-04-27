package main

func main() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}

	folder1 := &Folder{
		component: nil,
		name:      "Folder11",
	}

	folder1.Add(file1)
	folder1.Search("rose")

	folder2 := &Folder{
		component: nil,
		name:      "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Search("ali")
}
