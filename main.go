package main

import (
	"fmt"

	"github.com/nenitf/ead_go_alura_oo/clientes"
	"github.com/nenitf/ead_go_alura_oo/contas"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	fmt.Println(contaDoBruno)
}
