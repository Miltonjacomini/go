package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const qtdeMonitoramentos = 1
const delayMonitoramento = 5

func main() {

	exibeBoasVindas()

	for {

		exibeOpcoes()

		opcaoEscolhida := leOpcao()

		executaOpcaoLida(opcaoEscolhida)

	}
}

func exibeBoasVindas() {
	fmt.Println("")
	fmt.Println("Bem vindo ao programa")
	fmt.Println("Versão 1.0")
	fmt.Println("")
}

func exibeOpcoes() {
	fmt.Println("")
	fmt.Println("Opções do sistema: ")
	fmt.Println("")
	fmt.Println("1- Monitorar")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Para sair do programa")
	fmt.Println("")
}

func leOpcao() int {
	var opcaoEscolhida int
	fmt.Scan(&opcaoEscolhida)

	return opcaoEscolhida
}

func executaOpcaoLida(opcaoEscolhida int) {

	switch opcaoEscolhida {
	case 1:
		monitoramento()
	case 2:
		exibeLogs()
	case 0:
		saiDoPrograma()
	default:
		fmt.Println("Opção não identificada.")
		os.Exit(-1)
	}

}

func monitoramento() {

	fmt.Println("")
	fmt.Println("-- Monitorando API's :")

	sites := leSitesDoArquivo()

	for i := 0; i < qtdeMonitoramentos; i++ {

		fmt.Println("--")
		if qtdeMonitoramentos > 1 {
			fmt.Println("Fase: ", i)
		}

		for _, site := range sites {
			testaSiteERegistraLog(site)
		}

		fmt.Println("--")
		fmt.Println("")
		time.Sleep(delayMonitoramento * time.Second)
	}

}

func testaSiteERegistraLog(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao iniciar monitoramento::")
		fmt.Println(err)
		os.Exit(-1)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Sucesso: ", site)
	} else {
		fmt.Println("Erro ao acessar: ", site, " status: ", resp.StatusCode)
	}

	registraLog(site, resp.StatusCode)
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("../resources/sites.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo: ", err)
		os.Exit(-1)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func registraLog(site string, httpStatus int) {

	arquivo, err := os.OpenFile("../resources/log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao registrar log::")
		fmt.Println(err)
		os.Exit(-1)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - site: " + site +
		" - status: " + strconv.Itoa(httpStatus) + "\n")

	arquivo.Close()
}

func exibeLogs() {
	fmt.Println("-- ")
	fmt.Println("Exibindo logs...")

	arquivo, err := ioutil.ReadFile("../resources/log.txt")
	if err != nil {
		fmt.Println("Problema ao ler arquivo", err)
	}

	fmt.Println("")
	fmt.Println(string(arquivo))
	fmt.Println("--")
}

func saiDoPrograma() {
	fmt.Println("Saindo do programa")
	os.Exit(0)
}
