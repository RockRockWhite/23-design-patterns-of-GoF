package Prototype

import "fmt"

type Folder struct {
    name string
}

func (f *Folder) Print() {
    fmt.Println(f.name)
}

func (f *Folder) Clone() ICloneable {
    return &Folder{name: f.name + "_clone"}
}

func CreateFolder(name string) *Folder {
    return &Folder{name: name}
}
