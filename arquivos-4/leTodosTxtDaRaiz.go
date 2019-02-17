package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	fmt.Println("--")
	fmt.Println("Listando arquivos com filepath.Walk ::")

	diretoriosEArquivos, err := retornaDirEArquivos("../")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(diretoriosEArquivos)

	fmt.Println("----")
	fmt.Println("--Imprimindo arquivos txt da raiz ::")
	for index, path := range diretoriosEArquivos {
		if strings.Contains(path, ".txt") {
			fmt.Println("{{")
			fmt.Println("Arquivo: ", index)
			leArquivo(path)
			fmt.Println("}}")
		}
	}

	fmt.Println("--")
	fmt.Println("----")

}

func retornaDirEArquivos(rootString string) ([]string, error) {

	var files []string

	err := filepath.Walk(rootString, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func leArquivo(path string) {

	data, err := ioutil.ReadFile(path)

	fmt.Println("--")
	fmt.Println("")
	fmt.Println("localfile.data ::")
	fmt.Println("")

	if err != nil {
		fmt.Println("Erro ao ler arquivo localfile.data em resources", err)
	} else {
		fmt.Println(string(data))
	}
	fmt.Println("")
	fmt.Println("--")
}
