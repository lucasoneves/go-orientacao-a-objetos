package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	clienteBruno := clientes.Titular{Nome: "Bruno", CPF: "223344000", Profissao: "Cientista de Dados"}

	contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 1122, NumeroConta: 2323, Saldo: 2200}

	fmt.Println(contaDoBruno)

	contaDoBruno.Depositar(-12200)

	fmt.Println(contaDoBruno)

}
