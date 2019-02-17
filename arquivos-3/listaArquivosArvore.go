package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	fmt.Println("--")
	fmt.Println("Listando arquivos com filepath.Walk ::")

	err := filepath.Walk("../",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("--")
}
