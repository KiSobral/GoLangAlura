package contas

import (
	"GoLangAlura/bank/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       clientes.Cliente
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Saque(valor float64) {
	if valor <= 0 {
		fmt.Println("Erro!")
		fmt.Println("Valor inválido!")
		return
	}

	if !(valor <= c.saldo) {
		fmt.Println("Erro!")
		fmt.Println("Valor maior do que o saldo!")
		return
	}

	fmt.Println("saldo antes do saque:  R$", c.saldo)
	c.saldo -= valor
	fmt.Println("saldo depois do saque: R$", c.saldo)

	fmt.Println("Saque de", valor, " reais realizado com sucesso!")
}

func (c *ContaCorrente) Depositar(valor float64) {
	if valor <= 0 {
		fmt.Println("Erro!")
		fmt.Println("Valor inválido!")
		return
	}

	fmt.Println("saldo antes do depósito:  R$", c.saldo)
	c.saldo += valor
	fmt.Println("saldo depois do depósito: R$", c.saldo)

	fmt.Println("Depósito de", valor, " reais realizado com sucesso!")
}

func (c *ContaCorrente) Transferencia(destino *ContaCorrente, valor float64) {
	if valor <= 0 {
		fmt.Println("Erro!")
		fmt.Println("Valor inválido!")
		return
	}

	if !(valor <= c.saldo) {
		fmt.Println("Erro!")
		fmt.Println("Valor maior do que o saldo!")
		return
	}

	fmt.Println("saldo origem antes da transferência:  R$", c.saldo)
	fmt.Println("saldo destino antes da transferência:  R$", destino.saldo)
	c.saldo -= valor
	destino.saldo += valor
	fmt.Println("saldo origem depois do depósito: R$", c.saldo)
	fmt.Println("saldo destino depois do depósito: R$", destino.saldo)

	fmt.Println("Transferência de", valor, " reais entre as contas de", c.Titular.Nome, "e", destino.Titular.Nome, "realizada com sucesso!")
}

func (c *ContaCorrente) VisualizarSaldo() float64 {
	return c.saldo
}
