package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	pessoa1 := clientes.Titular{Nome: "Lucas", CPF: "22334455", Profissao: "Software Engineer"}
	contaPoupancaDoLucas := contas.ContaPoupanca{Titular: pessoa1, NumeroAgencia: 2234, NumeroConta: 2234}

	contaPoupancaDoLucas.Depositar(200)

	contaPoupancaDoLucas.Sacar(50)

	contaPoupancaDoLucas.ObterSaldo()

	fmt.Println(contaPoupancaDoLucas)
	fmt.Println(contaPoupancaDoLucas.ObterSaldo())
}
