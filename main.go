package main

import (
	"fmt"

	"github.com/nenitf/ead_go_alura_oo/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	fmt.Println(contaDaSilvia.Saldo)

	fmt.Println(contaDaSilvia.Sacar(400))
	fmt.Println(contaDaSilvia.Saldo)

	fmt.Println(contaDaSilvia.Saldo)
	status, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(status, valor)

	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 10}
	status2 := contaDaSilvia.Transferir(200, &contaDoGustavo)
	fmt.Println(status2)
}
