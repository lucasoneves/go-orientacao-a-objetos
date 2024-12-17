package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque efetuado!"
	}

	return "Saque recusado - saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito efetuado! saldo:", c.saldo
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

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
