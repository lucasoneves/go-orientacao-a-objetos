package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	pessoa1 := clientes.Titular{Nome: "Lucas", CPF: "22334455", Profissao: "Software Engineer"}
	contaPoupancaDoLucas := contas.ContaPoupanca{Titular: pessoa1, NumeroAgencia: 2234, NumeroConta: 2234}

	contaPoupancaDoLucas.Depositar(200)

	// contaPoupancaDoLucas.ObterSaldo()

	fmt.Println(contaPoupancaDoLucas)

	PagarBoleto(&contaPoupancaDoLucas, 100)

	fmt.Println(contaPoupancaDoLucas)
}
