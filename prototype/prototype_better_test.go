package prototype

import (
	"fmt"
	"testing"
)

func TestFileTree(t *testing.T) {
	f1 := &File{name: "1.txt"}
	f2 := &File{name: "2.txt"}
	f3 := &File{name: "3.txt"}

	dir1 := &Folder{
		name:       "dir1",
		lowerFiles: []FileTree{f1},
	}

	dir2 := &Folder{
		name:       "dir2",
		lowerFiles: []FileTree{
			f1,
			dir1,
			f2,
			f3,
		},
	}

	fmt.Println("dir2:")
	dir2.Tree("  ")

	dir2Copy := dir2.Clone()
	fmt.Println("dir2 copy:")
	dir2Copy.Tree("  ")
}