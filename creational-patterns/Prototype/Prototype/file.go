package Prototype

import "fmt"

type File struct {
    name string
}

func (f *File) Print() {
    fmt.Println(f.name)
}

func (f *File) Clone() ICloneable {
    return &File{name: f.name + "_clone"}
}

func CreateFile(name string) *File {
    return &File{name: name}
}
