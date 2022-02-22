package main

import "prototype/Prototype"

func main() {
    var proto Prototype.ICloneable
    var clone Prototype.ICloneable

    proto = Prototype.CreateFile("File")
    proto.(*Prototype.File).Print()
    clone = proto.Clone()
    clone.(*Prototype.File).Print()

    proto = Prototype.CreateFolder("Folder")
    proto.(*Prototype.Folder).Print()
    clone = proto.Clone()
    clone.(*Prototype.Folder).Print()
}
