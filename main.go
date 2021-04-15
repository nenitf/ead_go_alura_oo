package main

import (
	"fmt"

	"github.com/nenitf/ead_go_alura_oo/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(-100)

	fmt.Println(contaExemplo.ObterSaldo())
}
