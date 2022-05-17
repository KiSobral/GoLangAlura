package main

import (
	"GoLangAlura/bank/clientes"
	"GoLangAlura/bank/contas"
	"fmt"
)

func PagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Saque(valorDoBoleto)
}

type verificaConta interface {
	Saque(valor float64)
}

func main() {
	clienteDenis := clientes.Cliente{
		Nome:      "Denis",
		CPF:       "01234567899",
		Profissao: "Desenvolvedor Go",
	}
	contaDoDenis := contas.ContaPoupanca{
		Titular: clienteDenis, NumeroAgencia: 123,
		NumeroConta: 654321, Operacao: 1,
	}
	fmt.Println(contaDoDenis)
	fmt.Println(contaDoDenis.VisualizarSaldo())
	contaDoDenis.Depositar(500)
	fmt.Println(contaDoDenis.VisualizarSaldo())
	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.VisualizarSaldo())
}
