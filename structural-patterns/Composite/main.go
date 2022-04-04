package main

import "Composite/Composite"

func main() {
    file1 := &Composite.File{Name: "File1"}
    file2 := &Composite.File{Name: "File2"}
    file3 := &Composite.File{Name: "File3"}

    folder1 := &Composite.Folder{
        Name: "Folder1",
    }

    folder1.Add(file1)

    folder2 := &Composite.Folder{
        Name: "Folder2",
    }
    folder2.Add(file2)
    folder2.Add(file3)
    folder2.Add(folder1)

    folder2.Search("rose")
}
