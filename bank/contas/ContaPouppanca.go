package contas

import (
	"GoLangAlura/bank/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular                              clientes.Cliente
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Saque(valor float64) {
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

func (c *ContaPoupanca) Depositar(valor float64) {
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

func (c *ContaPoupanca) VisualizarSaldo() float64 {
	return c.saldo
}
