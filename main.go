package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	var titular string = "Lucas de Oliveira Neves"
	var agencia int = 589
	var contaCorrente int = 123456
	var saldo float64 = 500.95

	fmt.Println("Titular da conta:", titular)
	fmt.Println("Agência:", agencia)
	fmt.Println("Conta Corrente:", contaCorrente)
	fmt.Println("Saldo atual:", saldo)

}
