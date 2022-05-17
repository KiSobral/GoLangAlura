package main

import (
	"GoLangAlura/bank/clientes"
	"GoLangAlura/bank/contas"
	"fmt"
)

func main() {
	clienteBruno := clientes.Cliente{
		Nome:      "Bruno",
		CPF:       "01234567890",
		Profissao: "Desenvolvedor",
	}
	contaDoBruno := contas.ContaCorrente{
		Titular: clienteBruno, NumeroAgencia: 123,
		NumeroConta: 123456,
	}
	fmt.Println(contaDoBruno.VisualizarSaldo())
	contaDoBruno.Depositar(500)
	fmt.Println(contaDoBruno.VisualizarSaldo())
}
