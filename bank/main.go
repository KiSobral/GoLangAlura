package main

import "GoLangAlura/bank/contas"

func main() {
	conta1 := contas.ContaCorrente{
		Titular:       "Hugo",
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         530,
	}

	conta2 := contas.ContaCorrente{
		Titular:       "Maria",
		NumeroAgencia: 632,
		NumeroConta:   654321,
		Saldo:         360,
	}

	conta1.Transferencia(&conta2, 85)
}
