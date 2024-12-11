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

	contaDoLucas := ContaCorrente{titular: "Lucas", saldo: 300}
	contaDaCamila := ContaCorrente{"Camila", 222, 34434, 400}

	fmt.Println(contaDoLucas)
	fmt.Println(contaDaCamila)

}
