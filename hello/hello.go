package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	file, fileErr := os.Open("sites.txt")
	if fileErr != nil {
		fmt.Println("Ocorreu um erro:", fileErr)
		return sites
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		sites = append(sites, strings.TrimSpace(line))
	}

	file.Close()
	return sites
}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:",
			response.StatusCode)
		registerLog(site, false)
	}
}

func registerLog(site string, status bool) {
	message := time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		"- online: " + strconv.FormatBool(status) + "\n"

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(message)
	file.Close()
}
