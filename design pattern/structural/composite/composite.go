package composite

import "fmt"

type IComponent interface {
	ShowName()
}

type File struct {
	Name string
}

func (f File) ShowName() {
	fmt.Println(f.Name)
}

type Folder struct {
	Name  string
	Files []IComponent
}

func (f Folder) ShowName() {
	fmt.Println(f.Name)
	for _, e := range f.Files {
		fmt.Print("\t")
		e.ShowName()
	}
}

// func main() {
// 	file1 := &composite.File{"file1"}
// 	file2 := &composite.File{"file2"}
// 	file3 := &composite.File{"file3"}
// 	file4 := &composite.File{"file4"}
// 	folder1 := &composite.Folder{"folder1", []composite.IComponent{file1}}
// 	folder2 := &composite.Folder{"folder2", []composite.IComponent{folder1, file2, file3, file4}}

// 	folder1.ShowName()
// 	folder2.ShowName()

// }
