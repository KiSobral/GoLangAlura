package main

import (
	"fmt"
	"net/http"
	"os"
)

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
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://alura.com.br/"
	sites[2] = "https://caelum.com.br/"
	sites[3] = "https://aprender3.unb.br/"
	fmt.Println("Monitorando...")
	response, _ := http.Get(sites[3])

	if response.StatusCode == 200 {
		fmt.Println("Site:", sites[3], "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", sites[3], "está com problemas. Status Code:",
			response.StatusCode)
	}
}
