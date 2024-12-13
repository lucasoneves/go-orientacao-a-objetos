package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaDoLucas := contas.ContaCorrente{Titular: "Lucas", Saldo: 300}

	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 1000}
	fmt.Println(contaDoGustavo)
	fmt.Println(contaDoLucas)

	status := contaDoLucas.Transferencia(150, &contaDoGustavo)

	fmt.Println(status)

	fmt.Println(contaDoGustavo)
	fmt.Println(contaDoLucas)

}
