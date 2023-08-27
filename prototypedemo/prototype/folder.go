package prototype

import "fmt"

type Folder struct {
	Children []INode
	Name     string
}

/*Print is function*/
func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, i := range f.Children {
		i.Print(s + s)
	}
}

/*Clone is function*/
func (f *Folder) Clone() INode {
	cloneFolder := &Folder{Name: f.Name + "_Clone"}
	var tempChildren []INode
	for _, i := range f.Children {
		copyFile := i.Clone()
		tempChildren = append(tempChildren, copyFile)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
