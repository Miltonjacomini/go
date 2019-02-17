package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const pathOrigem = "/Users/miltonjneto/Documents/Workspace/Tino/Blog/blog/themes/hermit/exampleSite/content/posts"
const pathDestino = "content/posts"
const extensaoArquivo = ".md"

func main() {

	diretoriosEArquivos, err := retornaDirEArquivos(pathOrigem)

	if err != nil {
		log.Println(err)
	}

	for _, path := range diretoriosEArquivos {
		if strings.Contains(path, extensaoArquivo) {
			copiaArquivoDePara(path, pathDestino)
		}
	}
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

func extraiNomeArquivoDoPath(path string) string {
	split := strings.Split(path, "/")
	size := len(split)

	return split[size-1]
}

func copiaArquivoDePara(pathOrigem string, pathDestino string) {

	data, err := ioutil.ReadFile(pathOrigem)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo origem: ", err)
	}

	nomeArquivo := extraiNomeArquivoDoPath(pathOrigem)
	pathCompleto := pathDestino + "/" + nomeArquivo

	_, exists := os.Stat(pathDestino)
	if exists != nil {
		os.MkdirAll(pathDestino, os.ModePerm)
	}

	arquivo, err := os.OpenFile(pathCompleto, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao gravar arquivo destino: ", err)
		fmt.Println(pathDestino)
	}

	arquivo.WriteString(string(data))
	arquivo.Close()

}
