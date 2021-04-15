package main

import (
	"fmt"

	"github.com/nenitf/ead_go_alura_oo/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(-100)

	fmt.Println(contaExemplo.ObterSaldo())

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(55)

	fmt.Println(contaDoDenis.ObterSaldo())
}
