package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	//var contaDoGuilherme ContaCorrente = ContaCorrente{}
	// contaDoGuilherme := ContaCorrente{}

	// contaDoGuilherme := ContaCorrente{titular: "Guilherme",
	// 	numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	// contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	// fmt.Println(contaDoGuilherme)
	// fmt.Println(contaDaBruna)

	// var contaDaCris *ContaCorrente // pq asterisco?
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"

	// fmt.Println(contaDaCris) // ponteiro
	// fmt.Println(*contaDaCris) // valor
	// fmt.Println(&contaDaCris) // endereço de memória

	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaDoGuilherme2 := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	fmt.Println(contaDoGuilherme == contaDoGuilherme2) // true (compara valores)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"

	fmt.Println(contaDaCris == contaDaCris2)   // false (compara endereços de memória)
	fmt.Println(*contaDaCris == *contaDaCris2) // true (compara valores)
}
