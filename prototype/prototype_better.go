package prototype

import "fmt"

//FileTree 文件系统中的原型接口
type FileTree interface {
	Tree(indentation string)
	Clone() FileTree
}

type File struct {
	name string
}

func (f *File) Tree(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) Clone() FileTree {
	return &File{name: f.name + "_copy"}
}

type Folder struct {
	name string
	lowerFiles []FileTree
}

func (f *Folder) Tree(indentation string) {
	fmt.Println(indentation + f.name)
	for _, file := range f.lowerFiles {
		file.Tree(indentation + indentation)
	}
}

func (f *Folder) Clone() FileTree {
	folder := &Folder{name: f.name + "_copy"}
	for _, file := range f.lowerFiles {
		folder.lowerFiles = append(folder.lowerFiles, file.Clone())
	}
	return folder
}
