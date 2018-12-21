package main

import (
	"log"
	"myrepo/aqua/entities"
	"os"
	"path/filepath"
)

func Traverse(dir string) chan *entities.File {

	output := make(chan *entities.File, 1000)

	go func() {
		defer close(output)

		if dir == "" {
			dir = "/"
		}

		err := filepath.Walk(dir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if !info.IsDir() {
					output <- &entities.File{
						Name:     info.Name(),
						Path:     path,
						Size:     info.Size(),
						Modified: info.ModTime(),
					}
				}

				return nil
			})
		if err != nil {
			log.Println(err)
		}

	}()

	return output
}
