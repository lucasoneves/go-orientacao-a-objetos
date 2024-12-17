package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque efetuado!"
	}

	return "Saque recusado - Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito efetuado! Saldo:", c.Saldo
	}

	return "Valor do depósito menor que zero", c.Saldo
}

func (c *ContaCorrente) Transferencia(valorDaTransf float64, contaDestino *ContaCorrente) bool {
	if valorDaTransf < c.Saldo && valorDaTransf > 0 {
		c.Saldo -= valorDaTransf
		contaDestino.Depositar(valorDaTransf)
		return true
	} else {
		return false
	}
}
