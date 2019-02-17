package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	exibeBoasVindas()

	for {
		exibeOpcoes()
		comando := leOpcoes()

		switch comando {
		case 1:
			leArquivo()
		case 2:
			escreveArquivo()
		case 3:
			limpaArquivo()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		}
	}

}

func exibeBoasVindas() {
	fmt.Println("Bem vindo ao Read/Write")
	fmt.Println("Versão 1.0")
}

func exibeOpcoes() {
	fmt.Println("Bem vindo ao Read/Write")
	fmt.Println(" ")
	fmt.Println("1- Para ler o arquivo")
	fmt.Println("2- Escrever o arquivo")
	fmt.Println("3- Limpar o arquivo")
	fmt.Println("0- Para sair do programa \n:")
}

func leOpcoes() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func leArquivo() {

	data, err := ioutil.ReadFile("../resources/localfile.data")

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

func escreveArquivo() {

	arquivo, err := os.OpenFile("../resources/localfile.data", os.O_CREATE|os.O_RDWR, 0666) //os.O_APPEND -> concatena no arquivo existente
	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString("Criando um arquivo e inputando texto nele.. \n um teste de quebra de linha também.")
	arquivo.Close()

	fmt.Println("--")
	fmt.Println("")
	fmt.Println("Finalizado!!")
	fmt.Println("")
	fmt.Println("--")
}

func limpaArquivo() {

	fmt.Println("--")
	fmt.Println("")

	arquivo, err := os.OpenFile("../resources/localfile.data", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Erro ao limpar arquivo: ", err)
	}

	arquivo.Truncate(0)
	arquivo.Close()

	fmt.Println("Finalizado!!")
	fmt.Println("")
	fmt.Println("--")

}
