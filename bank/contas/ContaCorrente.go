package contas

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Saque(valor float64) {
	if valor <= 0 {
		fmt.Println("Erro!")
		fmt.Println("Valor inválido!")
		return
	}

	if !(valor <= c.Saldo) {
		fmt.Println("Erro!")
		fmt.Println("Valor maior do que o Saldo!")
		return
	}

	fmt.Println("Saldo antes do saque:  R$", c.Saldo)
	c.Saldo -= valor
	fmt.Println("Saldo depois do saque: R$", c.Saldo)

	fmt.Println("Saque de", valor, " reais realizado com sucesso!")
}

func (c *ContaCorrente) Depositar(valor float64) {
	if valor <= 0 {
		fmt.Println("Erro!")
		fmt.Println("Valor inválido!")
		return
	}

	fmt.Println("Saldo antes do depósito:  R$", c.Saldo)
	c.Saldo += valor
	fmt.Println("Saldo depois do depósito: R$", c.Saldo)

	fmt.Println("Depósito de", valor, " reais realizado com sucesso!")
}

func (c *ContaCorrente) Transferencia(destino *ContaCorrente, valor float64) {
	if valor <= 0 {
		fmt.Println("Erro!")
		fmt.Println("Valor inválido!")
		return
	}

	if !(valor <= c.Saldo) {
		fmt.Println("Erro!")
		fmt.Println("Valor maior do que o Saldo!")
		return
	}

	fmt.Println("Saldo origem antes da transferência:  R$", c.Saldo)
	fmt.Println("Saldo destino antes da transferência:  R$", destino.Saldo)
	c.Saldo -= valor
	destino.Saldo += valor
	fmt.Println("Saldo origem depois do depósito: R$", c.Saldo)
	fmt.Println("Saldo destino depois do depósito: R$", destino.Saldo)

	fmt.Println("Transferência de", valor, " reais entre as contas de", c.Titular, "e", destino.Titular, "realizada com sucesso!")
}
