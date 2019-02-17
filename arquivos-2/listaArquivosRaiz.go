package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	files, err := ioutil.ReadDir("../")

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	fmt.Println("--")
	fmt.Println("Listando arquivos com ioutil.ReadDir ::")
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Println("--")

}
