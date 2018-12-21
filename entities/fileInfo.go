package entities

import (
	"fmt"
	"time"
)

type File struct {
	Name     string `json:"file"`
	Path     string
	Size     int64 `json:"size"`
	Modified time.Time
}

func (f *File) Print() {
	fmt.Printf("name: %s \npath: [%s] \nsize: [%d] \nmodified: [%s]\n\n\n", f.Name, f.Path, f.Size, f.Modified)
}
