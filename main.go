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

	pessoa2 := clientes.Titular{Nome: "Camila", CPF: "23345591", Profissao: "Analista de RH"}
	contaPoupancaDoLucas := contas.ContaPoupanca{Titular: pessoa1, NumeroAgencia: 2234, NumeroConta: 2234}

	contaDaCamila := contas.ContaCorrente{Titular: pessoa2, NumeroAgencia: 2340, NumeroConta: 12}

	contaPoupancaDoLucas.Depositar(200)

	contaDaCamila.Depositar(12000)

	// contaPoupancaDoLucas.ObterSaldo()

	fmt.Println(contaPoupancaDoLucas)

	fmt.Println(contaDaCamila.ObterSaldo())

	PagarBoleto(&contaPoupancaDoLucas, 100)

	PagarBoleto(&contaDaCamila, 300)

	fmt.Println(contaPoupancaDoLucas)
	fmt.Println(contaDaCamila.ObterSaldo())
}
