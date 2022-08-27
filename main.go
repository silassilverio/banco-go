package main

import (
	"fmt"

	"github.com/Banco/clientes"
	"github.com/Banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	cliente1 := clientes.Titular{Nome: "Filip", CPF: "123.456.789-10", Profissao: "Desenvolvedor"}
	cliente2 := clientes.Titular{Nome: "Silas", CPF: "987.654.321-00", Profissao: "Desenvolvedor"}

	conta1 := contas.ContaCorrente{Titular: cliente1, NumeroAgencia: 123, NumeroConta: 6453}
	conta2 := contas.ContaCorrente{Titular: cliente2, NumeroAgencia: 321, NumeroConta: 9683}

	conta2.Depositar(1000)
	conta1.Depositar(100)

	conta2.Tranferir(500, &conta1)

	conta2.Sacar(50)
	PagarBoleto(&conta1, 100)

	fmt.Println(cliente1.Nome+" tem o saldo de:", conta1.ObterSaldo())
	fmt.Println(cliente2.Nome+" tem o saldo de:", conta2.ObterSaldo())
}
