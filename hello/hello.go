package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorateNumber = 2
const delay = 3

func main() {
	var command int
	for {
		command = menu()
		controlCommand(command)
	}
}

func menu() int {
	var command int
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do Programa")
	fmt.Scan(&command)
	return command
}

func controlCommand(command int) {
	switch command {
	case 1:
		monitorate()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

	fmt.Println()
	fmt.Println()
}

func monitorate() {
	sites := getSites()

	fmt.Println("Monitorando...")

	for i := 0; i < monitorateNumber; i++ {
		for j, site := range sites {
			fmt.Println("Testando site", j+1, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}

}

func getSites() []string {
	sites := []string{}
	sites = append(sites, "https://random-status-code.herokuapp.com/")
	sites = append(sites, "https://alura.com.br/")
	sites = append(sites, "https://caelum.com.br/")
	sites = append(sites, "https://aprender3.unb.br/")
	return sites
}

func testSite(site string) {
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:",
			response.StatusCode)
	}
}
