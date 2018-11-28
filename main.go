package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os" //fala com o sistema operacional
	"strconv"
	"strings"
	"time"
)

const monitoramento = 2
const delay = 3

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:

			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 3:
			fmt.Println("saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("valor invalido")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	var nome string = "Gabriel"

	var versao = 1.1

	fmt.Println("ola, meu caro ", nome)
	fmt.Println("este programa esta na versão ", versao)

}

func leComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido) //pega informação no terminal e joga na varriavel comando
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")
	fmt.Println("")
}

func iniciarMonitoramento() {
	fmt.Println("monitorando...")

	sites := leiaSites()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("posição", i, "site", site)

			testaSite(site)
		}
		fmt.Println("")
		//delay de 5 segundos
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
	fmt.Println("")
}

func leiaSites() []string {
	var sites []string

	arquivo, err := os.Open("Sites.txt")

	if err != nil {
		fmt.Println("Erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println("site atual", linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + site + " - online:" + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
