package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque efetuado!"
	}

	return "Saque recusado - Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito efetuado! Saldo:", c.saldo
	}

	return "Valor do depósito menor que zero", c.saldo
}

func (c *ContaCorrente) Transferencia(valorDaTransf float64, contaDestino *ContaCorrente) bool {
	if valorDaTransf < c.saldo && valorDaTransf > 0 {
		c.saldo -= valorDaTransf
		contaDestino.Depositar(valorDaTransf)
		return true
	} else {
		return false
	}
}

func main() {

	contaDoLucas := ContaCorrente{titular: "Lucas", saldo: 300}

	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 1000}
	fmt.Println(contaDoGustavo)
	fmt.Println(contaDoLucas)

	status := contaDoLucas.Transferencia(150, &contaDoGustavo)

	fmt.Println(status)

	fmt.Println(contaDoGustavo)
	fmt.Println(contaDoLucas)

}
